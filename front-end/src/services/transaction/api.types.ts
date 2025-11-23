export type Result<TData, TError> = [TData, null] | [null, TError];

export type Response<TData> = {
  message: string;
  error: string;
  error_fields: Record<string, unknown>;
  result: TData;
};

export type TransactionType = "DEBIT" | "CREDIT";

export type TransactionStatus = "SUCCESS" | "PENDING" | "FAILED";

export type IssueResponse = {
  timestamp: string;
  name: string;
  type: TransactionType;
  amount: number;
  status: TransactionStatus;
  description: string;
};

export type BalanceResponse = {
  balance: number;
};
