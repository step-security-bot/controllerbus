/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "config";

/** Placeholder is a config used for type assertions. */
export interface Placeholder {
}

function createBasePlaceholder(): Placeholder {
  return {};
}

export const Placeholder = {
  encode(_: Placeholder, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Placeholder {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePlaceholder();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  // encodeTransform encodes a source of message objects.
  // Transform<Placeholder, Uint8Array>
  async *encodeTransform(
    source: AsyncIterable<Placeholder | Placeholder[]> | Iterable<Placeholder | Placeholder[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [Placeholder.encode(p).finish()];
        }
      } else {
        yield* [Placeholder.encode(pkt).finish()];
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, Placeholder>
  async *decodeTransform(
    source: AsyncIterable<Uint8Array | Uint8Array[]> | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<Placeholder> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [Placeholder.decode(p)];
        }
      } else {
        yield* [Placeholder.decode(pkt)];
      }
    }
  },

  fromJSON(_: any): Placeholder {
    return {};
  },

  toJSON(_: Placeholder): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<Placeholder>, I>>(base?: I): Placeholder {
    return Placeholder.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Placeholder>, I>>(_: I): Placeholder {
    const message = createBasePlaceholder();
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends Array<infer U> ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends { $case: string } ? { [K in keyof Omit<T, "$case">]?: DeepPartial<T[K]> } & { $case: T["$case"] }
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
