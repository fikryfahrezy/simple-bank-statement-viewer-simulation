function noop() {}

export type UsePaginationOptions<TData extends object> = {
  data: TData[];
  currentPage: number;
  limit: number;
  handlePageChange?: (page: number) => void;
};

export function usePagination<TData extends object>({
  data,
  currentPage,
  limit,
  handlePageChange = noop,
}: UsePaginationOptions<TData>) {
  const totalPages = data.length;
  const lastPage = Math.min(Math.max(currentPage + 2, 5), totalPages);
  const firstPage = Math.max(1, lastPage - 4);

  const offset = (currentPage - 1) * limit;
  const pagedData = data.slice(offset, limit);

  const goToNextPage = (): void => {
    const newPage = Math.min(currentPage + 1, totalPages);
    handlePageChange(newPage);
  };

  const goToPreviousPage = (): void => {
    const newPage = Math.max(currentPage - 1, 1);
    handlePageChange(newPage);
  };

  return {
    pagedData,
    firstPage,
    lastPage,
    goToPreviousPage,
    goToNextPage,
  };
}
