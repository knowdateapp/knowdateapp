import { useQuery } from '@tanstack/react-query';
import { IApiError } from 'shared/api';
import { Workspace } from 'shared/model';
import { INotesResponse } from '../types';
import { queriesKeys } from './queries-keys.ts';

export const useNotes = (workspace: Workspace) => {
  const queryKey = queriesKeys.notes(workspace);
  const result = useQuery<INotesResponse, IApiError>({
    queryKey,
  });

  return { ...result, queryKey };
};
