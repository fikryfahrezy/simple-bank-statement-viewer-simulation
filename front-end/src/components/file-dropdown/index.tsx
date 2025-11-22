import { useEffect, useRef, useState } from "react";
import styles from "./styles.module.css";

function noop() {}

export type FileDropdownProps = Omit<React.ComponentProps<"input">, "file"> & {
  onChange?: (files: FileList) => void;
};

function preventDefaults(event: Event) {
  event.preventDefault();
  event.stopPropagation();
}

export function FileDropdown({
  className,
  children,
  onChange = noop,
  ...restProps
}: FileDropdownProps) {
  const dropzoneRef = useRef<HTMLDivElement>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [isHighlighted, setIsHighlighted] = useState(false);

  useEffect(() => {
    if (!dropzoneRef.current || !fileInputRef.current) {
      return;
    }

    const aborter = new AbortController();

    dropzoneRef.current.addEventListener("dragenter", preventDefaults, {
      signal: aborter.signal,
    });
    dropzoneRef.current.addEventListener("dragover", preventDefaults, {
      signal: aborter.signal,
    });
    dropzoneRef.current.addEventListener("dragleave", preventDefaults, {
      signal: aborter.signal,
    });
    dropzoneRef.current.addEventListener("drop", preventDefaults, {
      signal: aborter.signal,
    });

    // For preventing browser default
    document.body.addEventListener("dragenter", preventDefaults, {
      signal: aborter.signal,
    });
    document.body.addEventListener("dragover", preventDefaults, {
      signal: aborter.signal,
    });
    document.body.addEventListener("dragleave", preventDefaults, {
      signal: aborter.signal,
    });
    document.body.addEventListener("drop", preventDefaults, {
      signal: aborter.signal,
    });

    function addHighlight() {
      setIsHighlighted(true);
    }

    function removeHighlight() {
      setIsHighlighted(false);
    }

    dropzoneRef.current.addEventListener("dragenter", addHighlight, {
      signal: aborter.signal,
    });

    dropzoneRef.current.addEventListener("dragover", addHighlight, {
      signal: aborter.signal,
    });

    dropzoneRef.current.addEventListener("dragleave", removeHighlight, {
      signal: aborter.signal,
    });

    dropzoneRef.current.addEventListener("drop", removeHighlight, {
      signal: aborter.signal,
    });

    function handleDrop(event: DragEvent) {
      const { dataTransfer } = event;
      if (!dataTransfer) {
        return;
      }
      const { files } = dataTransfer;
      onChange(files);
    }

    // Handle dropped files
    dropzoneRef.current.addEventListener("drop", handleDrop, false);

    return () => {
      aborter.abort();
    };
  }, [onChange]);

  return (
    <div
      ref={dropzoneRef}
      className={`${styles.zone} ${isHighlighted ? styles.zoneHighlighted : ""} ${className}`}
    >
      {children}
      <input
        {...restProps}
        ref={fileInputRef}
        type="file"
        style={{
          display: "none",
        }}
      />
    </div>
  );
}
