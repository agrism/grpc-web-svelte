/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./categories_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.CategoryServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.CategoryServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.IndexRequest,
 *   !proto.proto.Categories>}
 */
const methodDescriptor_CategoryService_Index = new grpc.web.MethodDescriptor(
  '/proto.CategoryService/Index',
  grpc.web.MethodType.UNARY,
  proto.proto.IndexRequest,
  proto.proto.Categories,
  /**
   * @param {!proto.proto.IndexRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Categories.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.IndexRequest,
 *   !proto.proto.Categories>}
 */
const methodInfo_CategoryService_Index = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.Categories,
  /**
   * @param {!proto.proto.IndexRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.Categories.deserializeBinary
);


/**
 * @param {!proto.proto.IndexRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.Categories)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.Categories>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.CategoryServiceClient.prototype.index =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.CategoryService/Index',
      request,
      metadata || {},
      methodDescriptor_CategoryService_Index,
      callback);
};


/**
 * @param {!proto.proto.IndexRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.Categories>}
 *     Promise that resolves to the response
 */
proto.proto.CategoryServicePromiseClient.prototype.index =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.CategoryService/Index',
      request,
      metadata || {},
      methodDescriptor_CategoryService_Index);
};


module.exports = proto.proto;

