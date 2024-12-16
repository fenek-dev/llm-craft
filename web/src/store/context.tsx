import React, { createContext, useState, useContext, useEffect } from "react";

interface Element {
  name: string;
  emoji: string;
}

interface ElementsContextType {
  elements: Element[];
  addElement: (element: Element) => void;
  removeElement: (name: string) => void;
  clearElements: () => void;
}

const ElementsContext = createContext<ElementsContextType | undefined>(
  undefined
);

export const ElementsProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [elements, setElements] = useState<Element[]>(() => {
    const savedElements = localStorage.getItem("elements");

    if (savedElements) {
      const result = JSON.parse(savedElements);
      if (Array.isArray(result) && result.length > 0) {
        return result;
      }
    }
    return [];
  });

  useEffect(() => {
    localStorage.setItem("elements", JSON.stringify(elements));
  }, [elements]);

  useEffect(() => {
    if (elements.length === 0) {
      fetch("http://localhost:8080/start")
        .then((res) => res.json())
        .then((data) => {
          setElements(
            data.map((e: any) => ({
              name: e.result,
              emoji: e.emoji,
            }))
          );
        });
    }
  }, [elements.length]);

  const addElement = (element: Element) => {
    setElements([...elements, element]);
  };

  const removeElement = (name: string) => {
    setElements(elements.filter((element) => element.name !== name));
  };

  const clearElements = () => {
    setElements([]);
  };

  const value = {
    elements,
    addElement,
    removeElement,
    clearElements,
  };

  return (
    <ElementsContext.Provider value={value}>
      {children}
    </ElementsContext.Provider>
  );
};

export const useElements = () => {
  const context = useContext(ElementsContext);
  if (context === undefined) {
    throw new Error("useElements must be used within an ElementsProvider");
  }
  return context;
};
