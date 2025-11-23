import styles from "./styles.module.css";

export type CardProps = React.ComponentProps<"div">;

export function Card({ className = "", ...restProps }: CardProps) {
  return <div {...restProps} className={`${styles.card} ${className}`} />;
}
