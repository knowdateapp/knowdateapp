import { Workspace } from 'shared/model';

const defaultKey = 'cards' as const;
export const queriesKeys = {
  all: defaultKey,
  cards: (workspace: Workspace) => [`${workspace}/cards`, defaultKey, workspace] as const,
};
