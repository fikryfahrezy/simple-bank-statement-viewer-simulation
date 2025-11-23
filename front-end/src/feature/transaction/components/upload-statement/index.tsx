"use client";

import styles from "./styles.module.css";
import { Button } from "@/components/button";
import { FileDropzone } from "@/components/file-dropzone";
import { Plus } from "@/components/icons";
import { Modal } from "@/components/modal";
import { noopAsync } from "@/utils/noop";
import { useState } from "react";

export type UploadStatementProps = {
  loading?: boolean;
  onUpload?: (file: File) => Promise<void>;
};

export function UploadStatement({
  loading = false,
  onUpload = noopAsync,
}: UploadStatementProps) {
  const [modalOpen, setModalOpen] = useState(false);
  const [files, setFiles] = useState<FileList>();

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  const onSubmit = () => {
    if (!files) {
      alert("Please upload the statment file in CSV format");
      return;
    }

    const file = files[0];
    if (!file) {
      alert(
        "Uh oh! Something wrong with the statement file, please try again.",
      );
      return;
    }

    if (file.type !== "text/csv") {
      alert("Only Accept CSV format.");
      return;
    }

    onUpload(file).then(() => {
      closeModal();
    });
  };

  return (
    <>
      <Button onClick={openModal}>
        <Plus /> Upload Statement
      </Button>
      <Modal open={modalOpen}>
        <form onSubmit={onSubmit}>
          <FileDropzone
            multiple={false}
            onChange={setFiles}
            accept="text/csv"
            aria-label="File drop zone, click or press Enter/Space to select a file"
          >
            <p>
              {loading
                ? "Uploading..."
                : files
                  ? "Your Statement File is Ready to Upload."
                  : "Upload Your Transaction Statement in CSV Format."}
            </p>
          </FileDropzone>
          <div className={styles.modalFooter}>
            <Button variant="secondary" type="button" onClick={closeModal}>
              Cancel
            </Button>
            <Button variant="primary" type="submit">
              Upload
            </Button>
          </div>
        </form>
      </Modal>
    </>
  );
}
