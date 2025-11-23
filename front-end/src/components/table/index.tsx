import styles from "./styles.module.css";

export type TableProps = React.ComponentProps<"table">;

export function Table({ className = "", ...restProps }: TableProps) {
  return <table {...restProps} className={`${styles.table} ${className}`} />;
}

export type TableRowProps = React.ComponentProps<"tr">;

export function TableRow({ className = "", ...restProps }: TableRowProps) {
  return <tr {...restProps} className={`${className}`} />;
}

export type TableHeaderProps = React.ComponentProps<"thead">;

export function TableHeader({
  className = "",
  ...restProps
}: TableHeaderProps) {
  return (
    <thead {...restProps} className={`${styles.tableHead} ${className}`} />
  );
}

export type TableHeadProps = React.ComponentProps<"th">;

export function TableHead({ className = "", ...restProps }: TableHeadProps) {
  return <th {...restProps} className={`${styles.tableHead} ${className}`} />;
}

export type TableBodyProps = React.ComponentProps<"tbody">;

export function TableBody({ className = "", ...restProps }: TableBodyProps) {
  return <tbody {...restProps} className={`${className}`} />;
}

export type TableCellProps = React.ComponentProps<"td">;

export function TableCell({ className = "", ...restProps }: TableCellProps) {
  return <td {...restProps} className={`${styles.tableCell} ${className}`} />;
}
