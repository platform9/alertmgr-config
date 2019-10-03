package alertmgrcfg

import (
	"context"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func deleteSecret(c client.Client, ns string, secretName string) (bool, error) {
	sec := v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: ns,
		},
	}

	err := c.Delete(context.TODO(), &sec)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func checkSecretExists(c client.Client, ns string, secretName string) (bool, error) {
	key := types.NamespacedName{Name: secretName, Namespace: ns}
	var sec v1.Secret

	err := c.Get(context.TODO(), key, &sec)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func createSecret(c client.Client, obj *metav1.ObjectMeta, ns string, secretName string, kind string, data []byte) error {

	trueVar := true
	cfg := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: ns,
			Annotations: map[string]string{
				"created_by": "alertmgr-controller",
			},
		},
		Data: map[string][]byte{
			"alertmanager.yaml": data,
		},
	}

	if obj != nil {
		cfg.ObjectMeta.OwnerReferences = []metav1.OwnerReference{
			metav1.OwnerReference{
				APIVersion: monitoringv1.SchemeGroupVersion.String(),
				Name:       obj.GetName(),
				Kind:       kind,
				UID:        obj.GetUID(),
				Controller: &trueVar,
			},
		}
	}

	if err := c.Create(context.TODO(), cfg); err != nil {
		return err
	}

	return nil
}
