package server

import (
	"net/http"

	"github.com/containers/libpod/pkg/api/handlers"
	"github.com/containers/libpod/pkg/api/handlers/generic"
	"github.com/containers/libpod/pkg/api/handlers/libpod"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerImagesHandlers(r *mux.Router) error {
	// swagger:operation POST /images/create compat createImage
	// ---
	// tags:
	//  - images (compat)
	// summary: Create an image
	// description: Create an image by either pulling it from a registry or importing it.
	// produces:
	// - application/json
	// parameters:
	//  - in: query
	//    name: fromImage
	//    type: string
	//    description: needs description
	//  - in: query
	//    name: fromSrc
	//    type: string
	//    description: needs description
	//  - in: query
	//    name: tag
	//    type: string
	//    description: needs description
	//  - in: header
	//    name: X-Registry-Auth
	//    type: string
	//    description: A base64-encoded auth configuration.
	//  - in: body
	//    name: request
	//    schema:
	//      type: string
	//    description: Image content if fromSrc parameter was used
	// responses:
	//   200:
	//     $ref: "#/responses/ok"
	//   404:
	//     $ref: "#/responses/NoSuchImage"
	//   500:
	//     $ref: "#/responses/InternalError"
	r.Handle(VersionedPath("/images/create"), APIHandler(s.Context, generic.CreateImageFromImage)).Methods(http.MethodPost).Queries("fromImage", "{fromImage}")
	r.Handle(VersionedPath("/images/create"), APIHandler(s.Context, generic.CreateImageFromSrc)).Methods(http.MethodPost).Queries("fromSrc", "{fromSrc}")
	// swagger:operation GET /images/json compat listImages
	// ---
	// tags:
	//  - images (compat)
	// summary: List Images
	// description: Returns a list of images on the server. Note that it uses a different, smaller representation of an image than inspecting a single image.
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DockerImageSummary"
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/images/json"), APIHandler(s.Context, generic.GetImages)).Methods(http.MethodGet)
	// swagger:operation POST /images/load compat loadImage
	// ---
	// tags:
	//  - images (compat)
	// summary: Import image
	// description: Load a set of images and tags into a repository.
	// parameters:
	//  - in: query
	//    name: quiet
	//    type: boolean
	//    description: not supported
	//  - in: body
	//    name: request
	//    description: tarball of container image
	//    schema:
	//      type: string
	// produces:
	// - application/json
	// responses:
	//   200:
	//     description: no error
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/images/load"), APIHandler(s.Context, handlers.LoadImage)).Methods(http.MethodPost)
	// swagger:operation POST /images/prune compat pruneImages
	// ---
	// tags:
	//  - images (compat)
	// summary: Prune unused images
	// description: Remove images from local storage that are not being used by a container
	// parameters:
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//      filters to apply to image pruning, encoded as JSON (map[string][]string). Available filters:
	//        - `dangling=<boolean>` When set to `true` (or `1`), prune only
	//           unused *and* untagged images. When set to `false`
	//           (or `0`), all unused images are pruned.
	//        - `until=<string>` Prune images created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	//        - `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune images with (or without, in case `label!=...` is used) the specified labels.
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsImageDeleteResponse"
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/images/prune"), APIHandler(s.Context, generic.PruneImages)).Methods(http.MethodPost)
	// swagger:operation GET /images/search compat searchImages
	// ---
	// tags:
	//  - images (compat)
	// summary: Search images
	// description: Search registries for an image
	// parameters:
	//  - in: query
	//    name: term
	//    type: string
	//    description: term to search
	//  - in: query
	//    name: limit
	//    type: integer
	//    description: maximum number of results
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//        A JSON encoded value of the filters (a `map[string][]string`) to process on the images list. Available filters:
	//        - `is-automated=(true|false)`
	//        - `is-official=(true|false)`
	//        - `stars=<number>` Matches images that has at least 'number' stars.
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsSearchResponse"
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/images/search"), APIHandler(s.Context, handlers.SearchImages)).Methods(http.MethodGet)
	// swagger:operation DELETE /images/{name} compat removeImage
	// ---
	// tags:
	//  - images (compat)
	// summary: Remove Image
	// description: Delete an image from local storage
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: name or ID of image to delete
	//  - in: query
	//    name: force
	//    type: boolean
	//    description: remove the image even if used by containers or has other tags
	//  - in: query
	//    name: noprune
	//    type: boolean
	//    description: not supported. will be logged as an invalid parameter if enabled
	// produces:
	//  - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsImageDeleteResponse"
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   409:
	//     $ref: '#/responses/ConflictError'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/images/name"), APIHandler(s.Context, handlers.RemoveImage)).Methods(http.MethodDelete)
	// swagger:operation GET /images/{name}/get compat exportImage
	// ---
	// tags:
	//  - images (compat)
	// summary: Export an image
	// description: Export an image in tarball format
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	// produces:
	//  - application/json
	// responses:
	//   200:
	//     description: no error
	//     schema:
	//      type: string
	//      format: binary
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/images/{name}/get"), APIHandler(s.Context, generic.ExportImage)).Methods(http.MethodGet)
	// swagger:operation GET /images/{name}/history compat imageHistory
	// ---
	// tags:
	//  - images (compat)
	// summary: History of an image
	// description: Return parent layers of an image.
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsHistory"
	//   404:
	//     $ref: "#/responses/NoSuchImage"
	//   500:
	//     $ref: "#/responses/InternalError"
	r.Handle(VersionedPath("/images/{name}/history"), APIHandler(s.Context, handlers.HistoryImage)).Methods(http.MethodGet)
	// swagger:operation GET /images/{name}/json compat inspectImage
	// ---
	// tags:
	//  - images (compat)
	// summary: Inspect an image
	// description: Return low-level information about an image.
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsImageInspect"
	//   404:
	//     $ref: "#/responses/NoSuchImage"
	//   500:
	//     $ref: "#/responses/InternalError"
	r.Handle(VersionedPath("/images/{name}/json"), APIHandler(s.Context, generic.GetImage))
	// swagger:operation POST /images/{name}/tag compat tagImage
	// ---
	// tags:
	//  - images (compat)
	// summary: Tag an image
	// description: Tag an image so that it becomes part of a repository.
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: repo
	//    type: string
	//    description: the repository to tag in
	//  - in: query
	//    name: tag
	//    type: string
	//    description: the name of the new tag
	// produces:
	// - application/json
	// responses:
	//   201:
	//     description: no error
	//   400:
	//     $ref: '#/responses/BadParamError'
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   409:
	//     $ref: '#/responses/ConflictError'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/images/{name}/tag"), APIHandler(s.Context, handlers.TagImage)).Methods(http.MethodPost)
	// swagger:operation POST /commit/ compat commitContainer
	// ---
	// tags:
	//  - commit (compat)
	// summary: Create a new image from a container
	// parameters:
	//  - in: query
	//    name: container
	//    type: string
	//    description: the name or ID of a container
	//  - in: query
	//    name: repo
	//    type: string
	//    description: the repository name for the created image
	//  - in: query
	//    name: tag
	//    type: string
	//    description: tag name for the created image
	//  - in: query
	//    name: comment
	//    type: string
	//    description: commit message
	//  - in: query
	//    name: author
	//    type: string
	//    description: author of the image
	//  - in: query
	//    name: pause
	//    type: boolean
	//    description: pause the container before committing it
	//  - in: query
	//    name: changes
	//    type: string
	//    description: instructions to apply while committing in Dockerfile format
	// produces:
	// - application/json
	// responses:
	//   201:
	//     description: no error
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/commit"), APIHandler(s.Context, generic.CommitContainer)).Methods(http.MethodPost)

	// swagger:operation POST /build images buildImage
	// ---
	// tags:
	//  - images
	// summary: Create image
	// description: Build an image from the given Dockerfile(s)
	// parameters:
	//  - in: query
	//    name: dockerfile
	//    type: string
	//    default: Dockerfile
	//    description: |
	//      Path within the build context to the `Dockerfile`.
	//      This is ignored if remote is specified and points to an external `Dockerfile`.
	//  - in: query
	//    name: t
	//    type: string
	//    default: latest
	//    description: A name and optional tag to apply to the image in the `name:tag` format.
	//  - in: query
	//    name: extrahosts
	//    type: string
	//    default:
	//    description: |
	//      TBD Extra hosts to add to /etc/hosts
	//      (As of version 1.xx)
	//  - in: query
	//    name: remote
	//    type: string
	//    default:
	//    description: |
	//      A Git repository URI or HTTP/HTTPS context URI.
	//      If the URI points to a single text file, the file’s contents are placed
	//      into a file called Dockerfile and the image is built from that file. If
	//      the URI points to a tarball, the file is downloaded by the daemon and the
	//      contents therein used as the context for the build. If the URI points to a
	//      tarball and the dockerfile parameter is also specified, there must be a file
	//      with the corresponding path inside the tarball.
	//      (As of version 1.xx)
	//  - in: query
	//    name: q
	//    type: boolean
	//    default: false
	//    description: |
	//      Suppress verbose build output
	//  - in: query
	//    name: nocache
	//    type: boolean
	//    default: false
	//    description: |
	//      Do not use the cache when building the image
	//      (As of version 1.xx)
	//  - in: query
	//    name: cachefrom
	//    type: string
	//    default:
	//    description: |
	//      JSON array of images used to build cache resolution
	//      (As of version 1.xx)
	//  - in: query
	//    name: pull
	//    type: boolean
	//    default: false
	//    description: |
	//      Attempt to pull the image even if an older image exists locally
	//      (As of version 1.xx)
	//  - in: query
	//    name: rm
	//    type: boolean
	//    default: true
	//    description: |
	//      Remove intermediate containers after a successful build
	//      (As of version 1.xx)
	//  - in: query
	//    name: forcerm
	//    type: boolean
	//    default: false
	//    description: |
	//      Always remove intermediate containers, even upon failure
	//      (As of version 1.xx)
	//  - in: query
	//    name: memory
	//    type: integer
	//    description: |
	//      Memory is the upper limit (in bytes) on how much memory running containers can use
	//      (As of version 1.xx)
	//  - in: query
	//    name: memswap
	//    type: integer
	//    description: |
	//      MemorySwap limits the amount of memory and swap together
	//      (As of version 1.xx)
	//  - in: query
	//    name: cpushares
	//    type: integer
	//    description: |
	//      CPUShares (relative weight
	//      (As of version 1.xx)
	//  - in: query
	//    name: cpusetcpus
	//    type: string
	//    description: |
	//      CPUSetCPUs in which to allow execution (0-3, 0,1)
	//      (As of version 1.xx)
	//  - in: query
	//    name: cpuperiod
	//    type: integer
	//    description: |
	//      CPUPeriod limits the CPU CFS (Completely Fair Scheduler) period
	//      (As of version 1.xx)
	//  - in: query
	//    name: cpuquota
	//    type: integer
	//    description: |
	//      CPUQuota limits the CPU CFS (Completely Fair Scheduler) quota
	//      (As of version 1.xx)
	//  - in: query
	//    name: buildargs
	//    type: string
	//    default:
	//    description: |
	//      JSON map of string pairs denoting build-time variables.
	//      For example, the build argument `Foo` with the value of `bar` would be encoded in JSON as `["Foo":"bar"]`.
	//
	//      For example, buildargs={"Foo":"bar"}.
	//
	//      Note(s):
	//      * This should not be used to pass secrets.
	//      * The value of buildargs should be URI component encoded before being passed to the API.
	//
	//      (As of version 1.xx)
	//  - in: query
	//    name: shmsize
	//    type: integer
	//    default: 67108864
	//    description: |
	//      ShmSize is the "size" value to use when mounting an shmfs on the container's /dev/shm directory.
	//      Default is 64MB
	//      (As of version 1.xx)
	//  - in: query
	//    name: squash
	//    type: boolean
	//    default: false
	//    description: |
	//      Silently ignored.
	//      Squash the resulting images layers into a single layer
	//      (As of version 1.xx)
	//  - in: query
	//    name: labels
	//    type: string
	//    default:
	//    description: |
	//      JSON map of key, value pairs to set as labels on the new image
	//      (As of version 1.xx)
	//  - in: query
	//    name: networkmode
	//    type: string
	//    default: bridge
	//    description: |
	//      Sets the networking mode for the run commands during build.
	//      Supported standard values are:
	//        * `bridge` limited to containers within a single host, port mapping required for external access
	//        * `host` no isolation between host and containers on this network
	//        * `none` disable all networking for this container
	//        * container:<nameOrID> share networking with given container
	//        ---All other values are assumed to be a custom network's name
	//      (As of version 1.xx)
	//  - in: query
	//    name: platform
	//    type: string
	//    default:
	//    description: |
	//      Platform format os[/arch[/variant]]
	//      (As of version 1.xx)
	//  - in: query
	//    name: target
	//    type: string
	//    default:
	//    description: |
	//      Target build stage
	//      (As of version 1.xx)
	//  - in: query
	//    name: outputs
	//    type: string
	//    default:
	//    description: |
	//      output configuration TBD
	//      (As of version 1.xx)
	// produces:
	// - application/json
	// responses:
	//   200:
	//     description: OK (As of version 1.xx)
	//     schema:
	//       type: object
	//       required:
	//         - stream
	//       properties:
	//         stream:
	//           type: string
	//           description: output from build process
	//           example: |
	//             (build details...)
	//             Successfully built 8ba084515c724cbf90d447a63600c0a6
	//   400:
	//     $ref: "#/responses/BadParamError"
	//   500:
	//     $ref: "#/responses/InternalError"
	r.Handle(VersionedPath("/build"), APIHandler(s.Context, handlers.BuildImage)).Methods(http.MethodPost)
	/*
		libpod endpoints
	*/

	// swagger:operation POST /libpod/images/{name}/exists libpod libpodImageExists
	// ---
	// tags:
	//  - images
	// summary: Image exists
	// description: Check if image exists in local store
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   204:
	//     description: image exists
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/{name}/exists"), APIHandler(s.Context, libpod.ImageExists))
	r.Handle(VersionedPath("/libpod/images/{name}/tree"), APIHandler(s.Context, libpod.ImageTree))
	// swagger:operation GET /libpod/images/{name}/history libpod libpodImageHistory
	// ---
	// tags:
	//  - images
	// summary: History of an image
	// description: Return parent layers of an image.
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsHistory"
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/{name}/history"), APIHandler(s.Context, handlers.HistoryImage)).Methods(http.MethodGet)
	// swagger:operation GET /libpod/images/json libpod libpodListImages
	// ---
	// tags:
	//  - images
	// summary: List Images
	// description: Returns a list of images on the server
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DockerImageSummary"
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/json"), APIHandler(s.Context, libpod.GetImages)).Methods(http.MethodGet)
	// swagger:operation POST /libpod/images/load libpod libpodLoadImage
	// ---
	// tags:
	//  - images
	// summary: Import image
	// description: Load a set of images and tags into a repository.
	// parameters:
	//  - in: query
	//    name: quiet
	//    type: boolean
	//    description: not supported
	//  - in: body
	//    name: request
	//    description: tarball of container image
	//    required: true
	//    schema:
	//      type: string
	// produces:
	// - application/json
	// responses:
	//   200:
	//     description: no error
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/load"), APIHandler(s.Context, handlers.LoadImage)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/images/prune libpod libpodPruneImages
	// ---
	// tags:
	//  - images
	// summary: Prune unused images
	// description: Remove images that are not being used by a container
	// parameters:
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//      filters to apply to image pruning, encoded as JSON (map[string][]string). Available filters:
	//        - `dangling=<boolean>` When set to `true` (or `1`), prune only
	//           unused *and* untagged images. When set to `false`
	//           (or `0`), all unused images are pruned.
	//        - `until=<string>` Prune images created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	//        - `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune images with (or without, in case `label!=...` is used) the specified labels.
	//  - in: query
	//    name: all
	//    type: boolean
	//    description: prune all images
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsImageDeleteResponse"
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/prune"), APIHandler(s.Context, libpod.PruneImages)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/images/search libpod libpodSearchImages
	// ---
	// tags:
	//  - images
	// summary: Search images
	// description: Search registries for images
	// parameters:
	//  - in: query
	//    name: term
	//    type: string
	//    description: term to search
	//  - in: query
	//    name: limit
	//    type: integer
	//    description: maximum number of results
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//        A JSON encoded value of the filters (a `map[string][]string`) to process on the images list. Available filters:
	//        - `is-automated=(true|false)`
	//        - `is-official=(true|false)`
	//        - `stars=<number>` Matches images that has at least 'number' stars.
	// produces:
	// - application/json
	// responses:
	//   200:
	//      $ref: "#/responses/DocsSearchResponse"
	//   500:
	//      $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/search"), APIHandler(s.Context, handlers.SearchImages)).Methods(http.MethodGet)
	// swagger:operation DELETE /libpod/images/{name} libpod libpodRemoveImage
	// ---
	// tags:
	//  - images
	// summary: Remove Image
	// description: Delete an image from local store
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: name or ID of image to delete
	//  - in: query
	//    name: force
	//    type: boolean
	//    description: remove the image even if used by containers or has other tags
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsImageDeleteResponse"
	//   400:
	//     $ref: "#/responses/BadParamError"
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   409:
	//     $ref: '#/responses/ConflictError'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/name"), APIHandler(s.Context, handlers.RemoveImage)).Methods(http.MethodDelete)
	// swagger:operation GET /libpod/images/{name}/get libpod libpoodExportImage
	// ---
	// tags:
	//  - images
	// summary: Export an image
	// description: Export an image as a tarball
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: format
	//    type: string
	//    description: format for exported image
	//  - in: query
	//    name: compress
	//    type: boolean
	//    description: use compression on image
	// produces:
	// - application/json
	// responses:
	//   200:
	//     description: no error
	//     schema:
	//      type: string
	//      format: binary
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/{name}/get"), APIHandler(s.Context, libpod.ExportImage)).Methods(http.MethodGet)
	// swagger:operation GET /libpod/images/{name}/json libpod libpodInspectImage
	// ---
	// tags:
	//  - images
	// summary: Inspect an image
	// description: Obtain low-level information about an image
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   200:
	//     $ref: "#/responses/DocsLibpodInspectImageResponse"
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/{name}/json"), APIHandler(s.Context, libpod.GetImage))
	// swagger:operation POST /libpod/images/{name}/tag libpod libpodTagImage
	// ---
	// tags:
	//  - images
	// summary: Tag an image
	// description: Tag an image so that it becomes part of a repository.
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: repo
	//    type: string
	//    description: the repository to tag in
	//  - in: query
	//    name: tag
	//    type: string
	//    description: the name of the new tag
	// produces:
	// - application/json
	// responses:
	//   201:
	//     description: no error
	//   400:
	//     $ref: '#/responses/BadParamError'
	//   404:
	//     $ref: '#/responses/NoSuchImage'
	//   409:
	//     $ref: '#/responses/ConflictError'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/images/{name}/tag"), APIHandler(s.Context, handlers.TagImage)).Methods(http.MethodPost)

	return nil
}
