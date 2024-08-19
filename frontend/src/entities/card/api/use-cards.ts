import { useQuery } from '@tanstack/react-query';
import { IApiError } from 'shared/api';
import { Workspace } from 'shared/types';
import { CardsResponse } from '../types';
import { queriesKeys } from './queries-keys.ts';

export const useCards = (workspace: Workspace) => {
  const queryKey = queriesKeys.cards(workspace);
  return useQuery<CardsResponse, IApiError>({
    queryKey,
  });
};
