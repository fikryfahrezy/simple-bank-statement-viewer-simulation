"use client";

import styles from "./styles.module.css";
import { Button } from "@/components/button";
import { FileDropzone } from "@/components/file-dropzone";
import { File, Plus } from "@/components/icons";
import { Modal } from "@/components/modal";
import { noopAsync } from "@/utils/noop";
import { FormEvent, useState } from "react";

export type UploadStatementProps = {
  loading?: boolean;
  onUpload?: (file: File) => Promise<void>;
};

export function UploadStatement({
  loading = false,
  onUpload = noopAsync,
}: UploadStatementProps) {
  const [modalOpen, setModalOpen] = useState(false);
  const [files, setFiles] = useState<FileList | null>(null);

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  const onSubmit = (event: FormEvent) => {
    event.preventDefault();

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
      setFiles(null);
    });
  };

  return (
    <>
      <Button onClick={openModal}>
        <Plus width={16} height={16} /> Upload Statement
      </Button>
      <Modal open={modalOpen}>
        <form onSubmit={onSubmit} className={styles.form}>
          <FileDropzone
            multiple={false}
            onChange={setFiles}
            accept="text/csv"
            aria-label="File drop zone, click or press Enter/Space to select a file"
          >
            {loading ? (
              <p>Uploading...</p>
            ) : files ? (
              <div className={styles.fileReadyContainer}>
                <File />
                Your Statement File is Ready to Upload.
              </div>
            ) : (
              <p>Upload Your Transaction Statement in CSV Format.</p>
            )}
          </FileDropzone>
          <div className={styles.modalFooter}>
            <Button
              variant="secondary"
              type="button"
              className={styles.actionButton}
              onClick={closeModal}
            >
              Cancel
            </Button>
            <Button
              variant="primary"
              type="submit"
              className={styles.actionButton}
            >
              Upload
            </Button>
          </div>
        </form>
      </Modal>
    </>
  );
}
