// Ref: https://github.com/themesberg/flowbite-react/blob/38913e51536ecf1c8e924a6de571880cb91d2ea0/packages/ui/src/components/Pagination/Pagination.tsx
export function range(start: number, end: number): number[] {
  if (start >= end) {
    return [];
  }

  return Array(end - start + 1)
    .keys()
    .toArray()
    .map((key) => key + start);
}
