package ownerutil

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
)

// Owner is used to build an OwnerReference, and we need type and object metadata
type Owner interface {
	metav1.Object
	runtime.Object
	schema.ObjectKind
}

func IsOwnedBy(object metav1.Object, owner Owner) bool {
	for _, oref := range object.GetOwnerReferences() {
		if oref.UID == owner.GetUID() {
			return true
		}
	}
	return false
}

func IsOwnedByKind(object metav1.Object, ownerKind string) bool {
	for _, oref := range object.GetOwnerReferences() {
		if oref.Kind == ownerKind {
			return true
		}
	}
	return false
}

func GetOwnerByKind(object metav1.Object, ownerKind string) metav1.OwnerReference {
	for _, oref := range object.GetOwnerReferences() {
		if oref.Kind == ownerKind {
			return oref
		}
	}
	return metav1.OwnerReference{}
}

// AddNonBlockingOwner adds a nonblocking owner to the ownerref list.
func AddNonBlockingOwner(object metav1.Object, owner Owner) {
	// Most of the time we won't have TypeMeta on the object, so we infer it for types we know about
	if err := InferGroupVersionKind(owner); err != nil {
		log.Warn(err.Error())
	}
	blockOwnerDeletion := false
	isController := false

	ownerRefs := object.GetOwnerReferences()
	if ownerRefs == nil {
		ownerRefs = []metav1.OwnerReference{}
	}
	gvk := owner.GroupVersionKind()
	apiVersion, kind := gvk.ToAPIVersionAndKind()
	ownerRefs = append(ownerRefs, metav1.OwnerReference{
		APIVersion:         apiVersion,
		Kind:               kind,
		Name:               owner.GetName(),
		UID:                owner.GetUID(),
		BlockOwnerDeletion: &blockOwnerDeletion,
		Controller:         &isController,
	})
	object.SetOwnerReferences(ownerRefs)
}

// AddOwner adds an owner to the ownerref list.
func AddOwner(object metav1.Object, owner Owner, blockOwnerDeletion, isController bool) {
	// Most of the time we won't have TypeMeta on the object, so we infer it for types we know about
	if err := InferGroupVersionKind(owner); err != nil {
		log.Warn(err.Error())
	}

	ownerRefs := object.GetOwnerReferences()
	if ownerRefs == nil {
		ownerRefs = []metav1.OwnerReference{}
	}
	gvk := owner.GroupVersionKind()
	apiVersion, kind := gvk.ToAPIVersionAndKind()
	ownerRefs = append(ownerRefs, metav1.OwnerReference{
		APIVersion:         apiVersion,
		Kind:               kind,
		Name:               owner.GetName(),
		UID:                owner.GetUID(),
		BlockOwnerDeletion: &blockOwnerDeletion,
		Controller:         &isController,
	})
	object.SetOwnerReferences(ownerRefs)
}

// InferGroupVersionKind adds TypeMeta to an owner so that it can be written to an ownerref.
// TypeMeta is generally only known at serialization time, so we often won't know what GVK an owner has.
// For the types we know about, we can add the GVK of the apis that we're using the interact with the object.
func InferGroupVersionKind(obj runtime.Object) error {
	objectKind := obj.GetObjectKind()
	if !objectKind.GroupVersionKind().Empty() {
		// objectKind already has TypeMeta, no inference needed
		return nil
	}

	switch obj.(type) {
	case *corev1.ServiceAccount:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "ServiceAccount",
		})
	case *rbac.ClusterRole:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "rbac.authorization.k8s.io",
			Version: "v1",
			Kind:    "ClusterRole",
		})
	case *rbac.ClusterRoleBinding:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "rbac.authorization.k8s.io",
			Version: "v1",
			Kind:    "ClusterRoleBinding",
		})
	case *rbac.Role:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "rbac.authorization.k8s.io",
			Version: "v1",
			Kind:    "Role",
		})
	case *rbac.RoleBinding:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "rbac.authorization.k8s.io",
			Version: "v1",
			Kind:    "RoleBinding",
		})
	case *v1alpha1.ClusterServiceVersion:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   v1alpha1.GroupName,
			Version: v1alpha1.GroupVersion,
			Kind:    v1alpha1.ClusterServiceVersionKind,
		})
	case *v1alpha1.InstallPlan:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   v1alpha1.GroupName,
			Version: v1alpha1.GroupVersion,
			Kind:    v1alpha1.InstallPlanKind,
		})
	case *v1alpha1.Subscription:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   v1alpha1.GroupName,
			Version: v1alpha1.GroupVersion,
			Kind:    v1alpha1.SubscriptionKind,
		})
	case *v1alpha1.CatalogSource:
		objectKind.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   v1alpha1.GroupName,
			Version: v1alpha1.GroupVersion,
			Kind:    v1alpha1.CatalogSourceKind,
		})
	default:
		return fmt.Errorf("could not infer GVK for object: %#v, %#v", obj, objectKind)
	}
	return nil
}
