"use client";
import { Badge, BadgeProps, BadgeVariant } from "@/components/badge";
import styles from "./styles.module.css";
import { Button } from "@/components/button";
import {
  ChevronDown,
  ChevronLeft,
  ChevronRight,
  ChevronsUpDown,
  ChevronUp,
} from "@/components/icons";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/table";
import { SortOrder, useSorting } from "@/hooks/use-sorting";
import {
  TransactionStatus,
  TransactionType,
} from "@/services/transaction/api.types";
import { range } from "@/utils/array";
import { idTimeFormat } from "@/utils/date";
import { idrFormatter } from "@/utils/currency";

export type TransactionRow = {
  timestamp: string;
  name: string;
  type: TransactionType;
  amount: number;
  status: TransactionStatus;
  description: string;
};

export type TransactionTableProps = {
  rows: TransactionRow[];
  loading?: boolean;
  sortKey: keyof TransactionRow | null;
  sortOrder: SortOrder | null;
  onSortChange: (
    key: keyof TransactionRow | null,
    order: SortOrder | null,
  ) => void;
};

export function TransactionTable({
  rows,
  loading = false,
  sortKey,
  sortOrder,
  onSortChange,
}: TransactionTableProps) {
  const { sortedData, toggleOrder, order, key } = useSorting({
    data: rows,
    key: sortKey,
    order: sortOrder,
    onSortChange,
  });

  return (
    <div className={styles.tableContainer}>
      <Table>
        <TableHeader>
          <TableRow>
            <TransactionTableHead
              currentSortKey={key}
              sortKey="timestamp"
              sortOrder={order}
              toggleOrder={toggleOrder}
            >
              Timestamp
            </TransactionTableHead>
            <TransactionTableHead
              currentSortKey={key}
              sortKey="name"
              sortOrder={order}
              toggleOrder={toggleOrder}
            >
              Name
            </TransactionTableHead>
            <TransactionTableHead
              currentSortKey={key}
              sortKey="type"
              sortOrder={order}
              toggleOrder={toggleOrder}
            >
              Type
            </TransactionTableHead>
            <TransactionTableHead
              currentSortKey={key}
              sortKey="amount"
              sortOrder={order}
              toggleOrder={toggleOrder}
            >
              Amount
            </TransactionTableHead>
            <TransactionTableHead
              currentSortKey={key}
              sortKey="status"
              sortOrder={order}
              toggleOrder={toggleOrder}
            >
              Status
            </TransactionTableHead>
            <TransactionTableHead
              currentSortKey={key}
              sortKey="description"
              sortOrder={order}
              toggleOrder={toggleOrder}
            >
              Description
            </TransactionTableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {sortedData.map((row, rowIndex) => {
            return (
              <TableRow key={rowIndex}>
                <TableCell>{idTimeFormat(row.timestamp)}</TableCell>
                <TableCell>{row.name}</TableCell>
                <TableCell>
                  <Badge>{row.type}</Badge>
                </TableCell>
                <TableCell>{idrFormatter(row.amount)}</TableCell>
                <TableCell>
                  <StatusBadge status={row.status}>{row.status}</StatusBadge>
                </TableCell>
                <TableCell>{row.description}</TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
      {loading && <TransactionTableLoading />}
      {sortedData.length === 0 && <TransactionTableEmpty />}
    </div>
  );
}

export type TransactionTableHeadProps = {
  currentSortKey: keyof TransactionRow | null;
  sortKey: keyof TransactionRow | null;
  toggleOrder: (newKey: keyof TransactionRow | null) => void;
  sortOrder: SortOrder | null;
  children?: React.ReactNode;
};

export function TransactionTableHead({
  children,
  currentSortKey,
  sortKey,
  toggleOrder,
  sortOrder,
}: TransactionTableHeadProps) {
  const onSort = () => {
    toggleOrder(sortKey);
  };

  return (
    <TableHead>
      <Button variant="ghost" onClick={onSort}>
        {children}
        <SortIcon active={sortKey === currentSortKey} sortOrder={sortOrder} />
      </Button>
    </TableHead>
  );
}

export type SortIconProps = React.ComponentProps<"svg"> & {
  active?: boolean;
  sortOrder: SortOrder | null;
};

export function SortIcon({
  sortOrder,
  active = false,
  ...restProps
}: SortIconProps) {
  const Icon = (() => {
    if (!active) {
      return ChevronsUpDown;
    }

    switch (sortOrder) {
      case "ASC":
        return ChevronUp;
      case "DESC":
        return ChevronDown;
      default:
        return ChevronsUpDown;
    }
  })();

  return <Icon {...restProps} width={16} height={16} />;
}

const STATUS_BADGE_VARIANT: Record<TransactionStatus, BadgeVariant> = {
  FAILED: "error",
  PENDING: "warning",
  SUCCESS: "primary",
};

export type StatusBadgeProps = BadgeProps & {
  status: TransactionStatus;
};

export function StatusBadge({ status, ...restProps }: StatusBadgeProps) {
  return <Badge {...restProps} variant={STATUS_BADGE_VARIANT[status]} />;
}

export type TransactionTablePagination = {
  currentPage: number;
  firstPage: number;
  lastPage: number;
  totalPages: number;
  goToPage: (page: number) => void;
  goToNextPage: () => void;
  goToPreviousPage: () => void;
};

export function TransactionTablePagination({
  currentPage,
  firstPage,
  lastPage,
  totalPages,
  goToPage,
  goToNextPage,
  goToPreviousPage,
}: TransactionTablePagination) {
  return (
    <nav role="navigation" aria-label="pagination">
      <ul className={styles.paginationList}>
        <li>
          <Button
            variant="secondary"
            className={styles.pageButton}
            disabled={currentPage === 1}
            onClick={goToPreviousPage}
          >
            <ChevronLeft width={16} height={16} />
          </Button>
        </li>
        {range(firstPage, lastPage).map((page) => {
          return (
            <Button
              key={page}
              variant={page === currentPage ? "primary" : "secondary"}
              className={styles.pageButton}
              onClick={() => {
                goToPage(page);
              }}
            >
              {page}
            </Button>
          );
        })}
        <li>
          <Button
            variant="secondary"
            className={styles.pageButton}
            disabled={currentPage >= totalPages}
            onClick={goToNextPage}
          >
            <ChevronRight width={16} height={16} />
          </Button>
        </li>
      </ul>
    </nav>
  );
}

export function TransactionTableLoading() {
  return <div className={styles.messageState}>Loading</div>;
}

export function TransactionTableEmpty() {
  return <div className={styles.messageState}>No Data</div>;
}
