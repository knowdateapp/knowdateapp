import { useMutation } from '@tanstack/react-query';
import { AxiosResponse } from 'axios';
import { cardsController, CreateCardParams, Card } from 'entities/card';
import { IApiError } from 'shared/api';

export const useCreateCard = () => {
  return useMutation<AxiosResponse<Card>, IApiError, CreateCardParams>({
    mutationFn: cardsController.createCard,
  });
};
