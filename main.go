package main

import (
	"log"
	"os"
	"time"

	"github.com/georgekuruvillak/machinetester/controller"

	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	machine "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned"
	machineinformer "github.com/gardener/machine-controller-manager/pkg/client/informers/externalversions"
	ctrl "sigs.k8s.io/controller-runtime"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	log.Print("Shared Informer app started")
	stopCh := ctrl.SetupSignalHandler()
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		LogErrAndExit(err, "error generating kubeconfig")
	}
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		LogErrAndExit(err, "error generating clientset")
	}
	machineClient, err := machine.NewForConfig(config)
	if err != nil {
		LogErrAndExit(err, "error generating clientset")
	}

	mgr, err := manager.New(config, ctrl.Options{
		LeaderElection: false,
	})
	if err != nil {
		LogErrAndExit(err, "error creating manager")
	}

	scheme := mgr.GetScheme()
	if err := machinev1alpha1.AddToScheme(scheme); err != nil {
		LogErrAndExit(err, "error adding machine scheme")
	}

	factory := machineinformer.NewSharedInformerFactory(machineClient, 5*time.Second)
	machineTester := controller.NewMachineTester(k8sClient,
		machineClient,
		factory.Machine().V1alpha1(),
	)

	machineTester.Start(stopCh)
	if err := mgr.Start(stopCh); err != nil {
		LogErrAndExit(err, "exiting manager with err")
	}
}

// LogErrAndExit logs the given error with msg and keysAndValues and calls `os.Exit(1)`.
func LogErrAndExit(err error, msg string, keysAndValues ...interface{}) {
	ctrl.Log.Error(err, msg, keysAndValues...)
	os.Exit(1)
}
