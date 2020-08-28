/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package containerregistryv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/container-registry-go-sdk/containerregistryv1"
)

/**
 * This file contains an integration test for the containerregistryv1 package.
 *
 * Notes:
 *
 * Your configuration file should contain the following variables. The following example is for the dev registry, using staging IAM
CONTAINER_REGISTRY_URL=https://dev.icr.io
CONTAINER_REGISTRY_AUTH_TYPE=iam
CONTAINER_REGISTRY_AUTH_URL=https://iam.test.cloud.ibm.com/identity/token
CONTAINER_REGISTRY_APIKEY=[An IAM Apikey]
CONTAINER_REGISTRY_ACCOUNT_ID=[Your test account ID]
CONTAINER_REGISTRY_NAMESPACE=[Namespace name, to be created and deleted by the test, eg: ]jahsdk
CONTAINER_REGISTRY_DNS_NAME=dev.icr.io
CONTAINER_REGISTRY_SEED_IMAGE=[An existing namespace/repo:tag to copy in this test, eg: ]newjhart/busy:latest
CONTAINER_REGISTRY_SEED_DIGEST=[The digest of the seed image, eg: ]sha256:2131f09e4044327fd101ca1fd4043e6f3ad921ae7ee901e9142e6e36b354a907
 *
 * The integration test will automatically skip tests if the required config file is not available.
*/

var _ = Describe(`ContainerRegistryV1 Integration Tests`, func() {

	const externalConfigFile = "../container_registry_v1.env"

	var (
		err               error
		containerRegistry *containerregistryv1.ContainerRegistryV1
		serviceURL        string
		baseNamespace     string
		accountID         string
		registryDNSName   string
		seedImage         string
		seedDigest        string

		config map[string]string
	)

	// Globlal variables to hold link values
	var (
		namespaceLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(containerregistryv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}
			baseNamespace = config["NAMESPACE"]
			if baseNamespace == "" {
				Skip("Unable to load baseNamespace configuration property, skipping tests")
			}
			accountID = config["ACCOUNT_ID"]
			if accountID == "" {
				Skip("Unable to load accountID configuration property, skipping tests")
			}
			registryDNSName = config["DNS_NAME"]
			if registryDNSName == "" {
				Skip("Unable to load registryDNSName configuration property, skipping tests")
			}
			seedImage = config["SEED_IMAGE"]
			if seedImage == "" {
				Skip("Unable to load seedImage configuration property, skipping tests")
			}
			seedDigest = config["SEED_DIGEST"]
			if seedDigest == "" {
				Skip("Unable to load seedDigest configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			containerRegistryOptions := &containerregistryv1.ContainerRegistryV1Options{
				Account: core.StringPtr(accountID),
			}

			containerRegistry, err = containerregistryv1.NewContainerRegistryV1UsingExternalConfig(containerRegistryOptions)

			Expect(err).To(BeNil())
			Expect(containerRegistry).ToNot(BeNil())
			Expect(containerRegistry.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateNamespace - Create namespace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNamespace(createNamespaceOptions *CreateNamespaceOptions)`, func() {

			createNamespaceOptions := &containerregistryv1.CreateNamespaceOptions{
				Namespace: core.StringPtr(baseNamespace),
			}

			namespace, response, err := containerRegistry.CreateNamespace(createNamespaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Or(Equal(201), Equal(200)))
			Expect(namespace).ToNot(BeNil())

			namespaceLink = *namespace.Namespace

		})
	})

	Describe(`GetAuth - Get authorization options`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAuth(getAuthOptions *GetAuthOptions)`, func() {

			getAuthOptions := &containerregistryv1.GetAuthOptions{}

			authOptions, response, err := containerRegistry.GetAuth(getAuthOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(authOptions).ToNot(BeNil())

		})
	})

	Describe(`UpdateAuth - Update authorization options`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAuth(updateAuthOptions *UpdateAuthOptions)`, func() {

			updateAuthOptions := &containerregistryv1.UpdateAuthOptions{
				IamAuthz: core.BoolPtr(true),
			}

			response, err := containerRegistry.UpdateAuth(updateAuthOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`ListImages - List images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListImages(listImagesOptions *ListImagesOptions)`, func() {

			listImagesOptions := &containerregistryv1.ListImagesOptions{
				Namespace:            core.StringPtr(namespaceLink),
				IncludeIBM:           core.BoolPtr(true),
				IncludePrivate:       core.BoolPtr(true),
				IncludeManifestLists: core.BoolPtr(true),
				Vulnerabilities:      core.BoolPtr(true),
			}

			remoteApiImage, response, err := containerRegistry.ListImages(listImagesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(remoteApiImage).ToNot(BeNil())

		})
	})

	Describe(`BulkDeleteImages - Bulk delete images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`BulkDeleteImages(bulkDeleteImagesOptions *BulkDeleteImagesOptions)`, func() {

			bulkDeleteImagesOptions := &containerregistryv1.BulkDeleteImagesOptions{
				BulkDelete: []string{fmt.Sprintf("%s/%s/notexist:1", registryDNSName, namespaceLink)},
			}

			imageBulkDeleteResult, response, err := containerRegistry.BulkDeleteImages(bulkDeleteImagesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageBulkDeleteResult).ToNot(BeNil())

		})
	})

	Describe(`ListImageDigests - List images by digest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListImageDigests(listImageDigestsOptions *ListImageDigestsOptions)`, func() {

			listImageDigestsOptions := &containerregistryv1.ListImageDigestsOptions{
				ExcludeTagged: core.BoolPtr(false),
				ExcludeVa:     core.BoolPtr(false),
				IncludeIbm:    core.BoolPtr(false),
			}

			digestListImage, response, err := containerRegistry.ListImageDigests(listImageDigestsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(digestListImage).ToNot(BeNil())

		})
	})

	Describe(`TagImage - Create tag`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TagImage(tagImageOptions *TagImageOptions)`, func() {

			tagImageOptions := &containerregistryv1.TagImageOptions{
				Fromimage: core.StringPtr(fmt.Sprintf("%s/%s", registryDNSName, seedImage)),
				Toimage:   core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			response, err := containerRegistry.TagImage(tagImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`InspectImage - Inspect an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InspectImage(inspectImageOptions *InspectImageOptions)`, func() {

			inspectImageOptions := &containerregistryv1.InspectImageOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			imageInspection, response, err := containerRegistry.InspectImage(inspectImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageInspection).ToNot(BeNil())

		})
	})

	Describe(`GetImageManifest - Get image manifest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetImageManifest(getImageManifestOptions *GetImageManifestOptions)`, func() {

			getImageManifestOptions := &containerregistryv1.GetImageManifestOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			response, err := containerRegistry.GetImageManifest(getImageManifestOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`GetMessages - Get messages`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMessages(getMessagesOptions *GetMessagesOptions)`, func() {

			getMessagesOptions := &containerregistryv1.GetMessagesOptions{}

			result, response, err := containerRegistry.GetMessages(getMessagesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
	})

	Describe(`ListNamespaces - List namespaces`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNamespaces(listNamespacesOptions *ListNamespacesOptions)`, func() {

			listNamespacesOptions := &containerregistryv1.ListNamespacesOptions{}

			result, response, err := containerRegistry.ListNamespaces(listNamespacesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
	})

	Describe(`ListNamespaceDetails - Detailed namespace list`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNamespaceDetails(listNamespaceDetailsOptions *ListNamespaceDetailsOptions)`, func() {

			listNamespaceDetailsOptions := &containerregistryv1.ListNamespaceDetailsOptions{}

			namespaceDetail, response, err := containerRegistry.ListNamespaceDetails(listNamespaceDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespaceDetail).ToNot(BeNil())

		})
	})

	Describe(`AssignNamespace - Assign namespace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AssignNamespace(assignNamespaceOptions *AssignNamespaceOptions)`, func() {

			assignNamespaceOptions := &containerregistryv1.AssignNamespaceOptions{
				Namespace:          core.StringPtr(namespaceLink),
				XAuthResourceGroup: core.StringPtr("aea257aa3c636f5e88267c4fd70f2c1f"),
			}

			namespace, response, err := containerRegistry.AssignNamespace(assignNamespaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(namespace).ToNot(BeNil())

		})
	})

	Describe(`GetPlans - Get plans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPlans(getPlansOptions *GetPlansOptions)`, func() {

			getPlansOptions := &containerregistryv1.GetPlansOptions{}

			plan, response, err := containerRegistry.GetPlans(getPlansOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(plan).ToNot(BeNil())

		})
	})

	Describe(`UpdatePlans - Update plans`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePlans(updatePlansOptions *UpdatePlansOptions)`, func() {

			updatePlansOptions := &containerregistryv1.UpdatePlansOptions{
				Plan: core.StringPtr("Standard"),
			}

			response, err := containerRegistry.UpdatePlans(updatePlansOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))

		})
	})

	Describe(`GetQuota - Get quotas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetQuota(getQuotaOptions *GetQuotaOptions)`, func() {

			getQuotaOptions := &containerregistryv1.GetQuotaOptions{}

			quota, response, err := containerRegistry.GetQuota(getQuotaOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(quota).ToNot(BeNil())

		})
	})

	Describe(`UpdateQuota - Update quotas`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateQuota(updateQuotaOptions *UpdateQuotaOptions)`, func() {

			updateQuotaOptions := &containerregistryv1.UpdateQuotaOptions{
				StorageMegabytes: core.Int64Ptr(int64(26)),
				TrafficMegabytes: core.Int64Ptr(int64(480)),
			}

			response, err := containerRegistry.UpdateQuota(updateQuotaOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`ListRetentionPolicies - List retention policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRetentionPolicies(listRetentionPoliciesOptions *ListRetentionPoliciesOptions)`, func() {

			listRetentionPoliciesOptions := &containerregistryv1.ListRetentionPoliciesOptions{}

			mapStringRetentionPolicy, response, err := containerRegistry.ListRetentionPolicies(listRetentionPoliciesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringRetentionPolicy).ToNot(BeNil())

		})
	})

	Describe(`SetRetentionPolicy - Set retention policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetRetentionPolicy(setRetentionPolicyOptions *SetRetentionPolicyOptions)`, func() {

			setRetentionPolicyOptions := &containerregistryv1.SetRetentionPolicyOptions{
				ImagesPerRepo:  core.Int64Ptr(int64(10)),
				Namespace:      core.StringPtr(namespaceLink),
				RetainUntagged: core.BoolPtr(false),
			}

			response, err := containerRegistry.SetRetentionPolicy(setRetentionPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))

		})
	})

	Describe(`AnalyzeRetentionPolicy - Analyze retention policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AnalyzeRetentionPolicy(analyzeRetentionPolicyOptions *AnalyzeRetentionPolicyOptions)`, func() {

			analyzeRetentionPolicyOptions := &containerregistryv1.AnalyzeRetentionPolicyOptions{
				ImagesPerRepo:  core.Int64Ptr(int64(10)),
				Namespace:      core.StringPtr(namespaceLink),
				RetainUntagged: core.BoolPtr(false),
			}

			mapStringstring, response, err := containerRegistry.AnalyzeRetentionPolicy(analyzeRetentionPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringstring).ToNot(BeNil())

		})
	})

	Describe(`GetRetentionPolicy - Get retention policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRetentionPolicy(getRetentionPolicyOptions *GetRetentionPolicyOptions)`, func() {

			getRetentionPolicyOptions := &containerregistryv1.GetRetentionPolicyOptions{
				Namespace: core.StringPtr(namespaceLink),
			}

			retentionPolicy, response, err := containerRegistry.GetRetentionPolicy(getRetentionPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(retentionPolicy).ToNot(BeNil())

		})
	})

	Describe(`DeleteImageTag - Delete tag`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteImageTag(deleteImageTagOptions *DeleteImageTagOptions)`, func() {

			deleteImageTagOptions := &containerregistryv1.DeleteImageTagOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:1", registryDNSName, namespaceLink)),
			}

			imageDeleteResult, response, err := containerRegistry.DeleteImageTag(deleteImageTagOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDeleteResult).ToNot(BeNil())

		})
	})

	Describe(`DeleteImage - Delete image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteImage(deleteImageOptions *DeleteImageOptions)`, func() {

			deleteImageOptions := &containerregistryv1.DeleteImageOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest@%s", registryDNSName, namespaceLink, seedDigest)),
			}
			imageDeleteResult, response, err := containerRegistry.DeleteImage(deleteImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageDeleteResult).ToNot(BeNil())

		})
	})

	Describe(`ListDeletedImages - List deleted images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDeletedImages(listDeletedImagesOptions *ListDeletedImagesOptions)`, func() {

			listDeletedImagesOptions := &containerregistryv1.ListDeletedImagesOptions{
				Namespace: core.StringPtr(namespaceLink),
			}

			mapStringTrash, response, err := containerRegistry.ListDeletedImages(listDeletedImagesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mapStringTrash).ToNot(BeNil())

		})
	})

	Describe(`RestoreTags - Restore a digest and all associated tags`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestoreTags(restoreTagsOptions *RestoreTagsOptions)`, func() {

			restoreTagsOptions := &containerregistryv1.RestoreTagsOptions{
				Digest: core.StringPtr(fmt.Sprintf("%s/%s/sdktest@%s", registryDNSName, namespaceLink, seedDigest)),
			}

			restoreResult, response, err := containerRegistry.RestoreTags(restoreTagsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(restoreResult).ToNot(BeNil())

		})
	})

	Describe(`RestoreImage - Restore deleted image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RestoreImage(restoreImageOptions *RestoreImageOptions)`, func() {

			restoreImageOptions := &containerregistryv1.RestoreImageOptions{
				Image: core.StringPtr(fmt.Sprintf("%s/%s/sdktest:nope", registryDNSName, namespaceLink)),
			}

			response, err := containerRegistry.RestoreImage(restoreImageOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))

		})
	})

	Describe(`DeleteNamespace - Delete namespace`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNamespace(deleteNamespaceOptions *DeleteNamespaceOptions)`, func() {

			deleteNamespaceOptions := &containerregistryv1.DeleteNamespaceOptions{
				Namespace: core.StringPtr(namespaceLink),
			}

			response, err := containerRegistry.DeleteNamespace(deleteNamespaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
