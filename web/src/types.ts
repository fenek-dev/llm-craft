export interface ElementType {
  name: string;
  emoji: string;
}

export interface ElementWithPosition extends ElementWithBoundaries {
  id: string;
  x: number;
  y: number;
}

export interface ElementWithBoundaries extends ElementType {
  bounds: DOMRect;
}
