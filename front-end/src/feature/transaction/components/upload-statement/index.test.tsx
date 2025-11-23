import { expect, test } from "vitest";
import { render, screen, within } from "@/test-utils";
import { UploadStatement } from "./index";

test("Make sure the modal appears when the button clicked", async () => {
  const { user } = render(<UploadStatement />);

  const uploadButton = screen.getByRole("button", { name: "Upload Statement" });
  expect(uploadButton).toBeVisible();
  expect(screen.queryByRole("dialog")).not.toBeInTheDocument();

  await user.click(uploadButton);
  expect(screen.getByRole("dialog")).toBeVisible();
  expect(
    screen.getByText("Upload Your Transaction Statement in CSV Format."),
  ).toBeVisible();
});

test("Make sure the modal appears when the button clicked", async () => {
  const { user } = render(<UploadStatement />);

  const uploadButton = screen.getByRole("button", { name: "Upload Statement" });
  expect(uploadButton).toBeVisible();
  expect(screen.queryByRole("dialog")).not.toBeInTheDocument();

  await user.click(uploadButton);
  expect(screen.getByRole("dialog")).toBeVisible();

  const file = new File(
    ["1624507883, JOHN DOE, DEBIT, 250000, SUCCESS, restaurant"],
    "test.txt",
    { type: "text/csv" },
  );

  await user.upload(
    within(
      screen.getByLabelText(
        "File drop zone, click or press Enter/Space to select a file",
      ),
    ).getByLabelText("File Dropzone"),
    file,
  );

  expect(
    screen.getByText("Your Statement File is Ready to Upload."),
  ).toBeVisible();
});

test("Make sure loading appears on uploading", async () => {
  const { user } = render(<UploadStatement loading={true} />);

  const uploadButton = screen.getByRole("button", { name: "Upload Statement" });
  expect(uploadButton).toBeVisible();
  expect(screen.queryByRole("dialog")).not.toBeInTheDocument();

  await user.click(uploadButton);
  expect(screen.getByRole("dialog")).toBeVisible();

  expect(screen.getByText("Uploading...")).toBeVisible();
});
