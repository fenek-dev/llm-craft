import { Draggable } from "@/components/draggable";
import { Button } from "@/components/ui/button";
import { generateRandomId } from "@/lib/id";
import { useElements } from "@/store/context";
import { ElementWithPosition } from "@/types";
import Sidebar from "@/widgets/sidebar";
import { useState } from "react";

const checkOverlap = (rect1: DOMRect, rect2: DOMRect) => {
  return !(
    rect1.right < rect2.left ||
    rect1.left > rect2.right ||
    rect1.bottom < rect2.top ||
    rect1.top > rect2.bottom
  );
};

function Main() {
  const [elementsOnBoard, setElements] = useState<ElementWithPosition[]>([]);
  const { addElement, clearElements } = useElements();

  const onDrop = async (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    const element = JSON.parse(event.dataTransfer.getData("element"));
    const x = event.clientX - element.bounds.width / 2;
    const y = event.clientY - element.bounds.height / 2;

    const domRect2 = new DOMRect(
      x,
      y,
      element.bounds.width,
      element.bounds.height
    );

    const id = generateRandomId();

    setElements((v) => [...v, { ...element, x, y, id: id }]);

    for (const elementOnBoard of elementsOnBoard) {
      const domRect1 = new DOMRect(
        elementOnBoard.x,
        elementOnBoard.y,
        elementOnBoard.bounds.width,
        elementOnBoard.bounds.height
      );

      if ("id" in element && elementOnBoard.id === element.id) {
        continue;
      }

      if (checkOverlap(domRect1, domRect2)) {
        try {
          const response = await fetch(
            `http://localhost:8080/pair?first=${elementOnBoard.name}&second=${element.name}`
          );

          if (!response.ok) {
            throw new Error(response.statusText);
          }

          const res = await response.json();

          addElement({
            name: res.result,
            emoji: res.emoji,
          });

          setElements((v) => {
            const a = [
              ...v.filter(
                (el) =>
                  el.id !== element.id &&
                  el.id !== elementOnBoard.id &&
                  el.id !== id
              ),
              {
                bounds: domRect1,
                x: elementOnBoard.x,
                y: elementOnBoard.y,
                id: generateRandomId(),
                name: res.result,
                emoji: res.emoji,
              },
            ];

            return a;
          });
          return;
        } catch (error) {
          console.error(error);
          return;
        }
      }
    }
  };

  const onDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault(); // Necessary to allow dropping
  };

  const cleanup = (id: string) => () => {
    setElements(elementsOnBoard.filter((element) => element.id !== id));
  };

  return (
    <main className="w-screen h-screen flex">
      <div
        className="w-full h-full relative"
        onDrop={onDrop}
        onDragOver={onDragOver}
      >
        {elementsOnBoard.map((element, index) => (
          <Draggable
            element={element}
            key={index}
            className="fixed"
            style={{ top: element.y, left: element.x }}
            variant="outline"
            onDragEnd={cleanup(element.id)}
          />
        ))}
        <div className="absolute bottom-4 right-4 flex gap-4">
          <Button variant="outline" onClick={() => setElements([])}>
            Clear Board
          </Button>
          <Button variant="outline" onClick={clearElements}>
            Clear All Elements
          </Button>
        </div>
      </div>
      <Sidebar />
    </main>
  );
}

export default Main;
