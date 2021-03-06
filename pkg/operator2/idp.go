package operator2

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	configv1 "github.com/openshift/api/config/v1"
	osinv1 "github.com/openshift/api/osin/v1"
)

var (
	scheme  = runtime.NewScheme()
	codecs  = serializer.NewCodecFactory(scheme)
	encoder = codecs.LegacyCodec(osinv1.GroupVersion) // TODO I think there is a better way to do this
)

func init() {
	utilruntime.Must(osinv1.Install(scheme))
}

func convertProviderConfigToOsinBytes(providerConfig *configv1.IdentityProviderConfig, syncData *idpSyncData, i int) ([]byte, error) {
	const missingProviderFmt string = "type %s was specified, but its configuration is missing"

	var p runtime.Object

	switch providerConfig.Type {
	case configv1.IdentityProviderTypeBasicAuth:
		basicAuthConfig := providerConfig.BasicAuth
		if basicAuthConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.BasicAuthPasswordIdentityProvider{
			RemoteConnectionInfo: configv1.RemoteConnectionInfo{
				URL: basicAuthConfig.URL,
				CA:  syncData.AddConfigMap(i, basicAuthConfig.CA, corev1.ServiceAccountRootCAKey, true),
				CertInfo: configv1.CertInfo{
					CertFile: syncData.AddSecret(i, basicAuthConfig.TLSClientCert, corev1.TLSCertKey, true),
					KeyFile:  syncData.AddSecret(i, basicAuthConfig.TLSClientKey, corev1.TLSPrivateKeyKey, true),
				},
			},
		}

	case configv1.IdentityProviderTypeGitHub:
		githubConfig := providerConfig.GitHub
		if githubConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.GitHubIdentityProvider{
			ClientID:      githubConfig.ClientID,
			ClientSecret:  syncData.AddSecretStringSource(i, githubConfig.ClientSecret, configv1.ClientSecretKey, false),
			Organizations: githubConfig.Organizations,
			Hostname:      githubConfig.Hostname,
			CA:            syncData.AddConfigMap(i, githubConfig.CA, corev1.ServiceAccountRootCAKey, true),
		}

	case configv1.IdentityProviderTypeGitLab:
		gitlabConfig := providerConfig.GitLab
		if gitlabConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.GitLabIdentityProvider{
			CA:           syncData.AddConfigMap(i, gitlabConfig.CA, corev1.ServiceAccountRootCAKey, true),
			URL:          gitlabConfig.URL,
			ClientID:     gitlabConfig.ClientID,
			ClientSecret: syncData.AddSecretStringSource(i, gitlabConfig.ClientSecret, configv1.ClientSecretKey, false),
			Legacy:       new(bool), // we require OIDC for GitLab now
		}

	case configv1.IdentityProviderTypeGoogle:
		googleConfig := providerConfig.Google
		if googleConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.GoogleIdentityProvider{
			ClientID:     googleConfig.ClientID,
			ClientSecret: syncData.AddSecretStringSource(i, googleConfig.ClientSecret, configv1.ClientSecretKey, false),
			HostedDomain: googleConfig.HostedDomain,
		}

	case configv1.IdentityProviderTypeHTPasswd:
		if providerConfig.HTPasswd == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.HTPasswdPasswordIdentityProvider{
			File: syncData.AddSecret(i, providerConfig.HTPasswd.FileData, configv1.HTPasswdDataKey, false),
		}

	case configv1.IdentityProviderTypeKeystone:
		keystoneConfig := providerConfig.Keystone
		if keystoneConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.KeystonePasswordIdentityProvider{
			RemoteConnectionInfo: configv1.RemoteConnectionInfo{
				URL: keystoneConfig.URL,
				CA:  syncData.AddConfigMap(i, keystoneConfig.CA, corev1.ServiceAccountRootCAKey, true),
				CertInfo: configv1.CertInfo{
					CertFile: syncData.AddSecret(i, keystoneConfig.TLSClientCert, corev1.TLSCertKey, true),
					KeyFile:  syncData.AddSecret(i, keystoneConfig.TLSClientKey, corev1.TLSPrivateKeyKey, true),
				},
			},
			DomainName:          keystoneConfig.DomainName,
			UseKeystoneIdentity: true, // force use of keystone ID
		}

	case configv1.IdentityProviderTypeLDAP:
		ldapConfig := providerConfig.LDAP
		if ldapConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.LDAPPasswordIdentityProvider{
			URL:          ldapConfig.URL,
			BindDN:       ldapConfig.BindDN,
			BindPassword: syncData.AddSecretStringSource(i, ldapConfig.BindPassword, configv1.BindPasswordKey, true),
			Insecure:     ldapConfig.Insecure,
			CA:           syncData.AddConfigMap(i, ldapConfig.CA, corev1.ServiceAccountRootCAKey, true),
		}

	case configv1.IdentityProviderTypeOpenID:
		openIDConfig := providerConfig.OpenID
		if openIDConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.OpenIDIdentityProvider{
			CA:                       syncData.AddConfigMap(i, openIDConfig.CA, corev1.ServiceAccountRootCAKey, true),
			ClientID:                 openIDConfig.ClientID,
			ClientSecret:             syncData.AddSecretStringSource(i, openIDConfig.ClientSecret, configv1.ClientSecretKey, false),
			ExtraScopes:              openIDConfig.ExtraScopes,
			ExtraAuthorizeParameters: openIDConfig.ExtraAuthorizeParameters,
			URLs: osinv1.OpenIDURLs{
				Authorize: openIDConfig.URLs.Authorize,
				Token:     openIDConfig.URLs.Token,
				UserInfo:  openIDConfig.URLs.UserInfo,
			},
			Claims: osinv1.OpenIDClaims{
				// There is no longer a user-facing setting for ID as it is considered unsafe
				ID:                []string{configv1.UserIDClaim},
				PreferredUsername: openIDConfig.Claims.PreferredUsername,
				Name:              openIDConfig.Claims.Name,
				Email:             openIDConfig.Claims.Email,
			},
		}

	case configv1.IdentityProviderTypeRequestHeader:
		requestHeaderConfig := providerConfig.RequestHeader
		if requestHeaderConfig == nil {
			return nil, fmt.Errorf(missingProviderFmt, providerConfig.Type)
		}

		p = &osinv1.RequestHeaderIdentityProvider{
			LoginURL:                 requestHeaderConfig.LoginURL,
			ChallengeURL:             requestHeaderConfig.ChallengeURL,
			ClientCA:                 syncData.AddConfigMap(i, requestHeaderConfig.ClientCA, corev1.ServiceAccountRootCAKey, false),
			ClientCommonNames:        requestHeaderConfig.ClientCommonNames,
			Headers:                  requestHeaderConfig.Headers,
			PreferredUsernameHeaders: requestHeaderConfig.PreferredUsernameHeaders,
			NameHeaders:              requestHeaderConfig.NameHeaders,
			EmailHeaders:             requestHeaderConfig.EmailHeaders,
		}

	default:
		return nil, fmt.Errorf("the identity provider type '%s' is not supported", providerConfig.Type)
	} // switch

	return encodeOrDie(p), nil
}

func createDenyAllIdentityProvider() osinv1.IdentityProvider {
	return osinv1.IdentityProvider{
		Name:            "defaultDenyAll",
		UseAsChallenger: true,
		UseAsLogin:      true,
		MappingMethod:   "claim",
		Provider: runtime.RawExtension{
			Raw: encodeOrDie(&osinv1.DenyAllPasswordIdentityProvider{}),
		},
	}
}

func encodeOrDie(obj runtime.Object) []byte {
	bytes, err := runtime.Encode(encoder, obj)
	if err != nil {
		panic(err) // indicates static generated code is broken, unrecoverable
	}
	return bytes
}
