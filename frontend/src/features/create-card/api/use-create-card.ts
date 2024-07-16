import { useMutation } from '@tanstack/react-query';
import { AxiosResponse } from 'axios';
import { cardsController, ICard } from '../../../entities/card';
import { IApiError } from '../../../shared/api';
import { ICreateCardParams } from '../../../entities/card/api';

export const useCreateCard = () => {
  return useMutation<AxiosResponse<ICard>, IApiError, ICreateCardParams>({
    mutationFn: cardsController.createCard,
  });
};
