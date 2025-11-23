import { BalanceResponse, IssueResponse, Response, Result } from "./api.types";

export class APIError extends Error {
  code: string;
  errorFields: Record<string, string[]> | null;
  constructor(
    code: string,
    message: string,
    errorFields: Record<string, string[]> | null,
  ) {
    super(message);
    this.name = "APIError";
    this.code = code;
    this.errorFields = errorFields;
  }
}

export class TransactionService {
  private baseURL: string;
  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  async uploadStatement(formData: FormData): Promise<Result<null, Error>> {
    const url = `${this.baseURL}/upload`;
    try {
      const response = await fetch(url, {
        method: "POST",
        body: formData,
      });
      const resBody = (await response.json()) as Response<null>;
      if (!response.ok) {
        return [
          null,
          new APIError(resBody.error, resBody.message, resBody.error_fields),
        ];
      }
      return [resBody.result, null];
    } catch (error) {
      return [null, new Error(String(error))];
    }
  }

  async getIssues(): Promise<Result<IssueResponse[], Error>> {
    const url = `${this.baseURL}/issues`;
    try {
      const response = await fetch(url);
      const resBody = (await response.json()) as Response<IssueResponse[]>;
      if (!response.ok) {
        return [
          null,
          new APIError(resBody.error, resBody.message, resBody.error_fields),
        ];
      }
      return [resBody.result, null];
    } catch (error) {
      return [null, new Error(String(error))];
    }
  }

  async getBalance(): Promise<Result<BalanceResponse, Error>> {
    const url = `${this.baseURL}/balance`;
    try {
      const response = await fetch(url);
      const resBody = (await response.json()) as Response<BalanceResponse>;
      if (!response.ok) {
        return [
          null,
          new APIError(resBody.error, resBody.message, resBody.error_fields),
        ];
      }
      return [resBody.result, null];
    } catch (error) {
      return [null, new Error(String(error))];
    }
  }
}

export const TRANSACTION_SERVICE_BASE_URL = process.env
  .NEXT_PUBLIC_TRANSACTION_SERVICE_BASE_URL as string;

export const transactionService = new TransactionService(
  TRANSACTION_SERVICE_BASE_URL,
);
