import styles from "./styles.module.css";

export type BadgeVariant = "primary" | "warning" | "error";

export type BadgeProps = React.ComponentProps<"span"> & {
  variant?: BadgeVariant;
};

export const BADGE_VARIANTS: Record<BadgeVariant, string> = {
  primary: styles.badgePrimary,
  warning: styles.badgeWarning,
  error: styles.badgeError,
};

export function Badge({
  className,
  variant = "primary",
  ...restProps
}: BadgeProps) {
  const variantClass = BADGE_VARIANTS[variant];
  return (
    <span
      {...restProps}
      className={`${className} ${variantClass} ${styles.badge}`}
    />
  );
}
