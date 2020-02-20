package controller

import (
	"fmt"

	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	machine "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned"
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/client/informers/externalversions/machine/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	ctrl "sigs.k8s.io/controller-runtime"
)

type MachineTester struct {
	k8sClient                 *kubernetes.Clientset
	machineClient             *machine.Clientset
	machineInformer           v1alpha1.MachineInformer
	machineSetInformer        v1alpha1.MachineSetInformer
	machineDeploymentInformer v1alpha1.MachineDeploymentInformer
}

func NewMachineTester(client *kubernetes.Clientset,
	machineClient *machine.Clientset,
	informerFactory v1alpha1.Interface,
) *MachineTester {
	return &MachineTester{
		k8sClient:                 client,
		machineClient:             machineClient,
		machineInformer:           informerFactory.Machines(),
		machineSetInformer:        informerFactory.MachineSets(),
		machineDeploymentInformer: informerFactory.MachineDeployments(),
	}
}

func (m *MachineTester) Start(stopChan <-chan struct{}) error {
	machineInformer := m.machineInformer.Informer()
	machineSetInformer := m.machineSetInformer.Informer()
	machineDeploymentInformer := m.machineDeploymentInformer.Informer()
	go machineInformer.Run(stopChan)
	if !cache.WaitForCacheSync(stopChan, machineInformer.HasSynced) {
		return fmt.Errorf("Timed out waiting for caches to sync")
	}
	go machineSetInformer.Run(stopChan)
	if !cache.WaitForCacheSync(stopChan, machineSetInformer.HasSynced) {
		return fmt.Errorf("Timed out waiting for caches to sync")
	}
	go machineDeploymentInformer.Run(stopChan)
	if !cache.WaitForCacheSync(stopChan, machineDeploymentInformer.HasSynced) {
		return fmt.Errorf("Timed out waiting for caches to sync")
	}
	return nil
}

// Reconcile reconciles the <req>.
func (r *MachineTester) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// SetupWithManager sets up manager with a new controller and r as the reconcile.Reconciler
func (r *MachineTester) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&machinev1alpha1.Machine{}).
		For(&machinev1alpha1.MachineDeployment{}).
		For(&machinev1alpha1.MachineSet{}).
		Complete(r)
}
