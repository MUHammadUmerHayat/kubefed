/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package controller

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"
	"github.com/marun/federation-v2/pkg/controller/federatedconfigmap"
	"github.com/marun/federation-v2/pkg/controller/federatedconfigmapoverride"
	"github.com/marun/federation-v2/pkg/controller/federatedconfigmapplacement"
	"github.com/marun/federation-v2/pkg/controller/federateddeployment"
	"github.com/marun/federation-v2/pkg/controller/federateddeploymentoverride"
	"github.com/marun/federation-v2/pkg/controller/federateddeploymentplacement"
	"github.com/marun/federation-v2/pkg/controller/federatedjob"
	"github.com/marun/federation-v2/pkg/controller/federatednamespaceplacement"
	"github.com/marun/federation-v2/pkg/controller/federatedreplicaset"
	"github.com/marun/federation-v2/pkg/controller/federatedreplicasetoverride"
	"github.com/marun/federation-v2/pkg/controller/federatedreplicasetplacement"
	"github.com/marun/federation-v2/pkg/controller/federatedsecret"
	"github.com/marun/federation-v2/pkg/controller/federatedsecretoverride"
	"github.com/marun/federation-v2/pkg/controller/federatedsecretplacement"
	"github.com/marun/federation-v2/pkg/controller/propagatedversion"
	"github.com/marun/federation-v2/pkg/controller/sharedinformers"
	"k8s.io/client-go/rest"
)

func GetAllControllers(config *rest.Config) ([]controller.Controller, chan struct{}) {
	shutdown := make(chan struct{})
	si := sharedinformers.NewSharedInformers(config, shutdown)
	return []controller.Controller{
		federatedconfigmap.NewFederatedConfigMapController(config, si),
		federatedconfigmapoverride.NewFederatedConfigMapOverrideController(config, si),
		federatedconfigmapplacement.NewFederatedConfigMapPlacementController(config, si),
		federateddeployment.NewFederatedDeploymentController(config, si),
		federateddeploymentoverride.NewFederatedDeploymentOverrideController(config, si),
		federateddeploymentplacement.NewFederatedDeploymentPlacementController(config, si),
		federatedjob.NewFederatedJobController(config, si),
		federatednamespaceplacement.NewFederatedNamespacePlacementController(config, si),
		federatedreplicaset.NewFederatedReplicaSetController(config, si),
		federatedreplicasetoverride.NewFederatedReplicaSetOverrideController(config, si),
		federatedreplicasetplacement.NewFederatedReplicaSetPlacementController(config, si),
		federatedsecret.NewFederatedSecretController(config, si),
		federatedsecretoverride.NewFederatedSecretOverrideController(config, si),
		federatedsecretplacement.NewFederatedSecretPlacementController(config, si),
		propagatedversion.NewPropagatedVersionController(config, si),
	}, shutdown
}
