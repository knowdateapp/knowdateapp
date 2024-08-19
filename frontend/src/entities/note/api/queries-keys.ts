import { Workspace } from 'shared/types';

const defaultKey = 'notes' as const;

export const queriesKeys = {
  all: defaultKey,
  notes: (workspace: Workspace) => [`${workspace}/notes`, defaultKey, workspace] as const,
};
