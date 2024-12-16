import React from "react";
import { Button, ButtonProps } from "./ui/button";
import { ElementType, ElementWithBoundaries } from "@/types";

type DraggableProps = {
  element: ElementType;
} & ButtonProps &
  React.RefAttributes<HTMLButtonElement>;

export const Draggable = ({
  element,
  onDragStart,
  ...rest
}: DraggableProps) => {
  return (
    <Button
      draggable
      onDragStart={(event) => {
        event.dataTransfer.setData(
          "element",
          JSON.stringify({
            ...element,
            bounds: event.currentTarget.getBoundingClientRect(),
          } as ElementWithBoundaries)
        );
      }}
      className="w-min"
      variant="outline"
      onClick={() => null}
      {...rest}
    >
      {element.emoji} {element.name}
    </Button>
  );
};
