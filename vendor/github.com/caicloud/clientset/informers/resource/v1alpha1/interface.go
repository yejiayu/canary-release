/*
Copyright 2017 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// StorageServices returns a StorageServiceInformer.
	StorageServices() StorageServiceInformer
	// StorageTypes returns a StorageTypeInformer.
	StorageTypes() StorageTypeInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// StorageServices returns a StorageServiceInformer.
func (v *version) StorageServices() StorageServiceInformer {
	return &storageServiceInformer{factory: v.SharedInformerFactory}
}

// StorageTypes returns a StorageTypeInformer.
func (v *version) StorageTypes() StorageTypeInformer {
	return &storageTypeInformer{factory: v.SharedInformerFactory}
}
