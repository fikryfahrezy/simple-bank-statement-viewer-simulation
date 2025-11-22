export type Result<TData, TError> = [TData, null] | [null, TError];

export type Response<TData> = {
  message: string;
  error: string;
  error_fields: Record<string, unknown>;
  result: TData;
};

export type IssueResponse = {
  timestamp: string;
  name: string;
  type: string;
  amount: number;
  status: string;
  description: string;
};

export type BalanceResponse = {
  balance: number;
};
