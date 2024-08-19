import { useQuery } from '@tanstack/react-query';
import { IApiError } from 'shared/api';
import { Workspace } from 'shared/types';
import { NotesResponse } from '../types';
import { queriesKeys } from './queries-keys.ts';

export const useNotes = (workspace: Workspace) => {
  const queryKey = queriesKeys.notes(workspace);
  return useQuery<NotesResponse, IApiError>({
    queryKey,
  });
};
