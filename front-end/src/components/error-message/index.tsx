import styles from "./styles.module.css";

export type ErrorMessageProps = {
  children?: React.ReactNode;
};

export function ErrorMessage({ children }: ErrorMessageProps) {
  return <span className={styles.error}>{children}</span>;
}
