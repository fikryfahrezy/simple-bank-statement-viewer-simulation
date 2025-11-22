import { transactionService } from "@/services/transaction/api-sdk";
import {
  BalanceResponse,
  IssueResponse,
} from "@/services/transaction/api.types";
import { useState } from "react";

export type QueryState = "idle" | "loading" | "error";

export function useUploadStatement() {
  const [state, setState] = useState<QueryState>("idle");
  const [error, setError] = useState<Error | null>(null);
  const [response, setResponse] = useState<IssueResponse[] | null>(null);

  const request = (
    ...params: Parameters<typeof transactionService.uploadStatement>
  ) => {
    setState("loading");
    transactionService.uploadStatement(...params).then((response) => {
      const [result, error] = response;
      if (error) {
        setState("error");
        setError(error);
        return;
      }
      setState("idle");
      setResponse(result);
    });

    return response;
  };

  return { response, request, state, error };
}

export function useGetIssues() {
  const [state, setState] = useState<QueryState>("idle");
  const [error, setError] = useState<Error | null>(null);
  const [response, setResponse] = useState<IssueResponse[] | null>(null);

  const request = (
    ...params: Parameters<typeof transactionService.getIssues>
  ) => {
    setState("loading");
    transactionService.getIssues(...params).then((response) => {
      const [result, error] = response;
      if (error) {
        setState("error");
        setError(error);
        return;
      }
      setState("idle");
      setResponse(result);
    });

    return response;
  };

  return { response, request, state, error };
}

export function useGetBalance() {
  const [state, setState] = useState<QueryState>("idle");
  const [error, setError] = useState<Error | null>(null);
  const [response, setResponse] = useState<BalanceResponse | null>(null);

  const request = (
    ...params: Parameters<typeof transactionService.getBalance>
  ) => {
    setState("loading");
    transactionService.getBalance(...params).then((response) => {
      const [result, error] = response;
      if (error) {
        setState("error");
        setError(error);
        return;
      }
      setState("idle");
      setResponse(result);
    });

    return response;
  };

  return { response, request, state, error };
}
