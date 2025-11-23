export type SortOrder = "ASC" | "DESC";

export type UseSoringOptions<TData extends object> = {
  data: TData[];
  key: keyof TData | null;
  order: SortOrder | null;
  onSortChange: (key: keyof TData | null, order: SortOrder | null) => void;
};

export function useSorting<TData extends object>({
  data,
  key,
  order,
  onSortChange,
}: UseSoringOptions<TData>) {
  const sortedData =
    key === null || order === null
      ? data
      : data.toSorted((a, b) => {
          const strA = String(a[key]).toLowerCase();
          const strB = String(b[key]).toLowerCase();
          if (order === "DESC") {
            return strB.localeCompare(strA);
          }
          return strA.localeCompare(strB);
        });

  const toggleOrder = (newKey: keyof TData | null) => {
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

    onSortChange(nextKey, nextOrder);
  };

  return { sortedData, key, order, toggleOrder };
}
