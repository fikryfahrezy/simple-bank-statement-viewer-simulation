import styles from "./styles.module.css";

export type TableProps = React.ComponentProps<"table">;

export function Table({ className, ...restProps }: TableProps) {
  return <table {...restProps} className={`${className} ${styles.table}`} />;
}

export type TableRowProps = React.ComponentProps<"tr">;

export function TableRow({ className, ...restProps }: TableRowProps) {
  return <tr {...restProps} className={`${className} ${styles.tableRow}`} />;
}

export type TableHeaderProps = React.ComponentProps<"thead">;

export function TableHeader({ className, ...restProps }: TableHeaderProps) {
  return (
    <thead {...restProps} className={`${className} ${styles.tableHead}`} />
  );
}

export type TableHeadProps = React.ComponentProps<"th">;

export function TableHead({ className, ...restProps }: TableHeadProps) {
  return <th {...restProps} className={`${className} ${styles.tableHead}`} />;
}

export type TableBodyProps = React.ComponentProps<"tbody">;

export function TableBody({ className, ...restProps }: TableBodyProps) {
  return (
    <tbody {...restProps} className={`${className} ${styles.tableBody}`} />
  );
}

export type TableCellProps = React.ComponentProps<"td">;

export function TableCell({ className, ...restProps }: TableCellProps) {
  return <td {...restProps} className={`${className} ${styles.tableCell}`} />;
}
