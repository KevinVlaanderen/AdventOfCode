import {Tile} from '../model/tiles/Tile';

export class TileMap<V> extends Map<Tile, V> {
  private static compare(a: Tile, b: Tile) {
    return a.coordinates.q === b.coordinates.q && a.coordinates.r === b.coordinates.r;
  }

  has(key: Tile): boolean {
    for (const _key of this.keys()) {
      if (TileMap.compare(key, _key)) {
        return true;
      }
    }
    return false;
  }

  get(key: Tile): V | undefined {
    for (const [_key, value] of this.entries()) {
      if (TileMap.compare(key, _key)) {
        return value;
      }
    }
    return undefined;
  }

  set(key: Tile, value: V): this {
    for (const _key of this.keys()) {
      if (TileMap.compare(key, _key)) {
        super.set(_key, value);
        return this;
      }
    }
    super.set(key, value);
    return this;
  }

  delete(key: Tile): boolean {
    for (const _key of this.keys()) {
      if (TileMap.compare(key, _key)) {
        return super.delete(_key);
      }
    }
    return false;
  }
}
