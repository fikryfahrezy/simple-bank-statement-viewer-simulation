import styles from "./styles.module.css";
import { createPortal } from "react-dom";
import { Card } from "../card";
import { useEffect, useRef } from "react";

export type ModalProps = {
  open?: boolean;
  children?: React.ReactNode;
};

export function Modal({ open = false, children }: ModalProps) {
  const modalContentRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (open) {
      document.body.style.position = "fixed";
      document.body.style.top = `-${window.scrollY}`;
    } else {
      document.body.style.position = "";
      document.body.style.top = "";
    }
  }, [open]);

  if (!open) {
    return null;
  }

  return createPortal(
    <div className={styles.backdrop}>
      <Card ref={modalContentRef} className={styles.container}>
        {children}
      </Card>
    </div>,
    document.body,
  );
}
