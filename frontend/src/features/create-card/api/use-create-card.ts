import { useMutation } from '@tanstack/react-query';
import { AxiosResponse } from 'axios';
import { cardsController, CreateCardParams, CardEntity } from 'entities/card';
import { IApiError } from 'shared/api';

export const useCreateCard = () => {
  return useMutation<AxiosResponse<CardEntity>, IApiError, CreateCardParams>({
    mutationFn: cardsController.createCard,
  });
};
