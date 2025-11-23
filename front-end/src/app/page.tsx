"use client";
import styles from "./page.module.css";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { Card } from "@/components/card";
import {
  useGetBalance,
  useGetIssues,
  useUploadStatement,
} from "@/feature/transaction/queries/transaction-service";
import { usePagination } from "@/hooks/use-pagination";
import { useEffect } from "react";
import {
  TransactionTable,
  TransactionTablePagination,
} from "@/feature/transaction/components/transaction-table";
import { SortOrder } from "@/hooks/use-sorting";
import { IssueResponse } from "@/services/transaction/api.types";
import { idrFormatter } from "@/utils/currency";
import { ErrorMessage } from "@/components/error-message";
import { UploadStatement } from "@/feature/transaction/components/upload-statement";

export default function Page() {
  const {
    response: issuesResponse,
    request: getIssues,
    state: issuesState,
    error: issuesError,
  } = useGetIssues();
  const {
    response: balanceResponse,
    request: getBalance,
    state: balanceState,
    error: balanceError,
  } = useGetBalance();
  const {
    request: uploadStatement,
    state: uploadState,
    error: uploadError,
  } = useUploadStatement();

  const issues = issuesResponse ?? [];
  const balance = balanceResponse?.balance || 0;

  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const page = searchParams.get("page");
  const currentPage = Number(page) || 1;

  const limit = searchParams.get("limit");
  const currentLimit = Number(limit) || 10;

  const sortKey = searchParams.get("sort_key") as keyof IssueResponse | null;
  let sortOrder = searchParams.get("sort_order") as SortOrder | null;
  if (sortOrder !== "ASC" && sortOrder !== "DESC") {
    sortOrder = null;
  }

  // Get a new searchParams string by merging the current
  // searchParams with a provided key/value pair
  const createQueryString = (keyVal: [string, string | null][]) => {
    const params = new URLSearchParams(searchParams.toString());
    keyVal.forEach(([key, value]) => {
      if (!value) {
        params.delete(key);
      } else {
        params.set(key, value);
      }
    });

    return params.toString();
  };

  const onPageChange = (page: number) => {
    const newQueryString = createQueryString([["page", String(page)]]);
    router.push(pathname + "?" + newQueryString);
  };

  const onSortChange = (
    key: keyof IssueResponse | null,
    order: SortOrder | null,
  ) => {
    const sortKey = key === null ? null : String(key);
    const sortOrder = order === null ? null : String(order);
    const newQueryString = createQueryString([
      ["sort_key", sortKey],
      ["sort_order", sortOrder],
    ]);
    router.push(pathname + "?" + newQueryString);
  };

  const {
    pagedData,
    firstPage,
    goToPage,
    goToNextPage,
    goToPreviousPage,
    lastPage,
    totalPages,
  } = usePagination({
    data: issues,
    currentPage,
    limit: currentLimit,
    handlePageChange: onPageChange,
  });

  const onUploadStatement = async (file: File) => {
    const formData = new FormData();
    formData.append("file", file);

    await uploadStatement(formData);
    await Promise.all([getIssues(), getBalance()]);
    onPageChange(totalPages);
  };

  useEffect(() => {
    getIssues();
    getBalance();
  }, [getIssues, getBalance]);

  return (
    <main>
      <Card>
        <div className={styles.header}>
          <div>
            <p className={styles.balance}>
              {balanceState === "loading"
                ? "Loading..."
                : idrFormatter(balance)}
            </p>
            {balanceError && (
              <ErrorMessage>{balanceError.message}</ErrorMessage>
            )}
          </div>
          <UploadStatement
            loading={uploadState === "loading"}
            onUpload={onUploadStatement}
          />
          {uploadError && <ErrorMessage>{uploadError.message}</ErrorMessage>}
        </div>
        <TransactionTable
          rows={pagedData}
          sortKey={sortKey}
          sortOrder={sortOrder}
          onSortChange={onSortChange}
          loading={issuesState === "loading"}
        />
        {issuesError && <ErrorMessage>{issuesError.message}</ErrorMessage>}
        <div className={styles.footer}>
          <p>
            Showing {currentPage} of {totalPages}
          </p>
          <TransactionTablePagination
            currentPage={currentPage}
            firstPage={firstPage}
            lastPage={lastPage}
            goToPage={goToPage}
            goToNextPage={goToNextPage}
            goToPreviousPage={goToPreviousPage}
            totalPages={totalPages}
          />
        </div>
      </Card>
    </main>
  );
}
