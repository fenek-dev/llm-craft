import { Button } from "@/components/ui/button";
import { useElements } from "@/store/context";
import { ElementWithBoundaries } from "@/types";

function Sidebar() {
  const { elements } = useElements();

  const onDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault(); // Necessary to allow dropping
  };

  return (
    <div
      className="w-96 h-screen border-l-2 border-stone-900 p-4 z-10"
      onDragOver={onDragOver}
    >
      <div className="flex flex-wrap gap-2">
        {elements.map((element, index) => (
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
            key={index}
            onClick={() => null}
          >
            {element.emoji} {element.name}
          </Button>
        ))}
      </div>
    </div>
  );
}

export default Sidebar;
