import { QueryClient } from '@tanstack/react-query';
import { apiClient } from '../api';

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      queryFn: ({ queryKey, meta }) =>
        apiClient.get(queryKey[0] as string, { params: meta }).then((response) => response.data),
    },
  },
});
