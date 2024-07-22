import { Workspace } from 'shared/model';

const defaultKey = 'notes' as const;

export const queriesKeys = {
  all: defaultKey,
  notes: (workspace: Workspace) => [`${workspace}/notes`, defaultKey, workspace] as const,
};
