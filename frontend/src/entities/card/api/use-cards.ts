import { Workspace } from '../../../shared/model';
import { queriesKeys } from '../../note/api';
import { useQuery } from '@tanstack/react-query';
import { ICardsResponse } from '../types';
import { IApiError } from '../../../shared/api';

export const useCards = (workspace: Workspace) => {
  const queryKey = queriesKeys.cards(workspace);
  const result = useQuery<ICardsResponse, IApiError>({
    queryKey,
  });

  // TODO: Зачем тут queryKey отдается?
  return { ...result, queryKey };
};
