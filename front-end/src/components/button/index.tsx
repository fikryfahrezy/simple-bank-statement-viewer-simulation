import styles from "./styles.module.css";

export type ButtonVariant = "primary" | "secondary" | "ghost";

export type ButtonProps<TAs extends React.ElementType = "button"> =
  React.ComponentProps<TAs> & {
    as?: TAs;
    variant?: ButtonVariant;
  };

export const BUTTON_VARIANTS: Record<ButtonVariant, string> = {
  primary: styles.buttonPrimary,
  secondary: styles.buttonSecondary,
  ghost: styles.buttonGhost,
};

export function Button<TAs extends React.ElementType = "button">({
  as: asProp,
  variant = "primary",
  className = "",
  ...restProps
}: ButtonProps<TAs>) {
  const Comp = asProp || "button";
  const variantClass = BUTTON_VARIANTS[variant as ButtonVariant];
  return (
    <Comp
      {...restProps}
      className={`${styles.button} ${variantClass} ${className}`}
    />
  );
}
