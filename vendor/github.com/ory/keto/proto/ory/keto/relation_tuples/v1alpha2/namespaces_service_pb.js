// source: ory/keto/relation_tuples/v1alpha2/namespaces_service.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

goog.exportSymbol('proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest', null, global);
goog.exportSymbol('proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse', null, global);
goog.exportSymbol('proto.ory.keto.relation_tuples.v1alpha2.Namespace', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.displayName = 'proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.repeatedFields_, null);
};
goog.inherits(proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.displayName = 'proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ory.keto.relation_tuples.v1alpha2.Namespace, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.relation_tuples.v1alpha2.Namespace.displayName = 'proto.ory.keto.relation_tuples.v1alpha2.Namespace';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest;
  return proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    namespacesList: jspb.Message.toObjectList(msg.getNamespacesList(),
    proto.ory.keto.relation_tuples.v1alpha2.Namespace.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse;
  return proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ory.keto.relation_tuples.v1alpha2.Namespace;
      reader.readMessage(value,proto.ory.keto.relation_tuples.v1alpha2.Namespace.deserializeBinaryFromReader);
      msg.addNamespaces(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNamespacesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.ory.keto.relation_tuples.v1alpha2.Namespace.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Namespace namespaces = 1;
 * @return {!Array<!proto.ory.keto.relation_tuples.v1alpha2.Namespace>}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.prototype.getNamespacesList = function() {
  return /** @type{!Array<!proto.ory.keto.relation_tuples.v1alpha2.Namespace>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ory.keto.relation_tuples.v1alpha2.Namespace, 1));
};


/**
 * @param {!Array<!proto.ory.keto.relation_tuples.v1alpha2.Namespace>} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse} returns this
*/
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.prototype.setNamespacesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.Namespace=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.Namespace}
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.prototype.addNamespaces = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.ory.keto.relation_tuples.v1alpha2.Namespace, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.prototype.clearNamespacesList = function() {
  return this.setNamespacesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.relation_tuples.v1alpha2.Namespace.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.Namespace} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.Namespace}
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.relation_tuples.v1alpha2.Namespace;
  return proto.ory.keto.relation_tuples.v1alpha2.Namespace.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.Namespace} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.Namespace}
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.relation_tuples.v1alpha2.Namespace.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.Namespace} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.Namespace} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.Namespace.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.ory.keto.relation_tuples.v1alpha2);
