import { Workspace } from '../../../shared/model';

export interface ICreateCardParams {
  workspace: Workspace;
  question: string;
  answer: string;
}
