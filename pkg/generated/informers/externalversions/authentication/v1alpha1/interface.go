// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/cluster-authentication-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AuthenticationOperatorConfigs returns a AuthenticationOperatorConfigInformer.
	AuthenticationOperatorConfigs() AuthenticationOperatorConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AuthenticationOperatorConfigs returns a AuthenticationOperatorConfigInformer.
func (v *version) AuthenticationOperatorConfigs() AuthenticationOperatorConfigInformer {
	return &authenticationOperatorConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
