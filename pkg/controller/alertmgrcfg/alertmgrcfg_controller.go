package alertmgrcfg

import (
	"context"
	"io/ioutil"
	"os"

	monitoringv1alpha1 "github.com/platform9/alertmgr-config/pkg/apis/monitoring/v1alpha1"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_alertmgrcfg")

const (
	configDir = "/etc/promplus"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new AlertMgrCfg Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileAlertMgrCfg{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("alertmgrcfg-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource AlertMgrCfg
	err = c.Watch(&source.Kind{Type: &monitoringv1alpha1.AlertMgrCfg{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner AlertMgrCfg
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &monitoringv1alpha1.AlertMgrCfg{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileAlertMgrCfg implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileAlertMgrCfg{}

// ReconcileAlertMgrCfg reconciles a AlertMgrCfg object
type ReconcileAlertMgrCfg struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a AlertMgrCfg object and makes changes based on the state read
// and what is in the AlertMgrCfg.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileAlertMgrCfg) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling AlertMgrCfg")

	// Fetch the AlertMgrCfg instance
	amc := &monitoringv1alpha1.AlertMgrCfg{}
	err := r.client.Get(context.TODO(), request.NamespacedName, amc)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	log.V(2).Info("syncing alert manager config: %s, type: %s", amc.Name, amc.Spec.Type)
	var alertManagerName string
	for key, val := range amc.Labels {
		log.V(2).Info("Labels: %s %s", key, val)
		if key == "alertmanager" {
			alertManagerName = val
			break
		}
	}
	log.V(2).Info("Alert manager label: %s", alertManagerName)

	if alertManagerName == "" {
		log.Info("Alert manager label missing in alertmanager config: %s", amc.Name)
		return reconcile.Result{}, nil
	}

	file, err := os.Open(configDir + "/alertmanager.yaml")
	if err != nil {
		log.Error(err, "Failed to open alert manager config file")
		return reconcile.Result{}, os.ErrInvalid
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err, "Failed to read alert manager config file")
		return reconcile.Result{}, os.ErrInvalid
	}

	var acfg alertConfig
	yaml.Unmarshal(data, &acfg)

	err = formatReceiver(amc, &acfg)
	if err != nil {
		log.Error(err, "Failed to format receiver for: %s", amc.Spec.Type)
		return reconcile.Result{}, err
	}

	var amcList *monitoringv1alpha1.AlertMgrCfgList
	err = r.client.List(context.TODO(), &client.ListOptions{}, amcList)

	if err != nil {
		log.Error(err, "Failed to get list of alert manager config objects ")
		return reconcile.Result{}, err
	}
	for _, amcItr := range amcList.Items {
		log.V(2).Info("Name: %s, ns: %s", amcItr.Name, amcItr.Namespace)

		for key, val := range amcItr.Labels {
			log.Info("Labels: %s %s", key, val)
			if key == "alertmanager" && val == alertManagerName {
				if amc.Name == amcItr.Name {
					log.V(2).Info("Ignoring current amc object: %s", amcItr.Name)
					continue
				}
				log.V(2).Info("Formatting receiver for: %s", amcItr.Name)
				err = formatReceiver(&amcItr, &acfg)
				if err != nil {
					log.Error(err, "Failed to format receiver for: %s", amc.Spec.Type)
					return reconcile.Result{}, err
				}
			}
		}
	}

	secretName := "alertmanager-" + alertManagerName

	exists, _ := checkSecretExists(r.client, amc.Namespace, secretName)
	if exists {
		log.Info("Secret for alertmanager: %s exists deleting it", amc.Name)
		_, err = deleteSecret(r.client, amc.Namespace, secretName)
		if err != nil {
			log.Error(err, "Failed to delete secret: %s", secretName)
			return reconcile.Result{}, err
		}
	}

	data, err = yaml.Marshal(&acfg)
	if err != nil {
		log.Error(err, "Failed to marshal alert mgr secret ")
		return reconcile.Result{}, err
	}

	err = createSecret(r.client, amc.ObjectMeta, secretName, "", data)
	if err != nil {
		return reconcile.Result{}, err
	}
	log.Info("Created secret: %s for %s", secretName, amc.Name)
	return reconcile.Result{}, nil
}
