import { useState } from "react";

export type SortOrder = "ASC" | "DESC";

export type UseSoringOptions<TData extends object> = {
  data: TData[];
};

export function useSorting<TData extends object>({
  data,
}: UseSoringOptions<TData>) {
  const [key, setKey] = useState<keyof TData | null>(null);
  const [order, setOrder] = useState<SortOrder | null>(null);

  const sortedData =
    key === null || order === null
      ? data
      : data.toSorted((a, b) => {
          return String(a[key])
            .toLowerCase()
            .localeCompare(String(b[key]).toLowerCase());
        });

  const toggleOrder = (newKey: keyof TData) => {
    let nextKey: keyof TData | null = newKey;
    let nextOrder: SortOrder | null = null;

    if (order === null) {
      nextOrder = "DESC";
    } else if (order === "DESC") {
      nextOrder = "ASC";
    } else {
      nextOrder = null;
      nextKey = null;
    }

    setKey(nextKey);
    setOrder(nextOrder);
  };

  return { sortedData, key, order, toggleOrder };
}
