/*
Copyright 2021 The Crossplane Authors.

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
package user

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_user", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "user"

		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
	})
}
