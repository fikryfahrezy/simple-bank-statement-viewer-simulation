import { ChangeEvent, useRef, useState } from "react";
import styles from "./styles.module.css";
import { noop } from "@/utils/noop";

export type FileDropzoneProps = Omit<
  React.ComponentProps<"input">,
  "file" | "onChange"
> & {
  files?: FileList;
  onChange?: (files: FileList) => void;
};

export function FileDropzone({
  className = "",
  children,
  onChange = noop,
  ...restProps
}: FileDropzoneProps) {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isDragging, setIsDragging] = useState(false);

  const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragging(true);
  };

  const handleDragLeave = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragging(false);
  };

  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragging(false);

    const files = event.dataTransfer?.files;
    if (!files) {
      return;
    }

    onChange(files);
  };

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (!files) {
      return;
    }

    onChange(files);
  };

  const onZoneClicked = () => {
    fileInputRef.current?.click();
  };

  const onZoneKeydown = (event: React.KeyboardEvent<HTMLDivElement>) => {
    if (event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      fileInputRef.current?.click();
    }
  };

  const style: React.CSSProperties = {
    borderColor: isDragging ? "darkgreen" : "#ccc",
    backgroundColor: isDragging ? "#f0fff0" : "#fff",
  };

  return (
    <div
      className={`${styles.zone} ${isDragging ? styles.zoneHighlighted : ""} ${className}`}
      tabIndex={0}
      role="button"
      aria-label={restProps["aria-label"]}
      onDragOver={handleDragOver}
      onDragLeave={handleDragLeave}
      onDrop={handleDrop}
      onClick={onZoneClicked}
      onKeyDown={onZoneKeydown}
    >
      <input
        {...restProps}
        ref={fileInputRef}
        type="file"
        style={{ display: "none" }}
        onChange={handleFileChange}
        aria-label="File Dropzone"
      />
      {children}
    </div>
  );
}
