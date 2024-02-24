# Руководство для разработчиков

## Git Flow

### Branches and commits

При разработке нового функционала нужно отводить ветку от ветки `main`. Новая ветка должна называться по
шаблону `feature/<feature_name>`. Где `<feature_name>` описывает добавляемый функционал или исправление.

```shell
git checkout main
git pull
git checkout -b feature/gif-flow
```

Работа в ветке ведется в свободной форме. Вы можете делать столько коммитов, сколько вам нужно для надежной фиксации
добавляемых изменений.

При составлении текста коммита нужно следовать
руководству [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0-beta.2/). Текст коммита должен
содержать в заголовке понятное и емкое описание проделанной работы. При желании подробное описание можно добавить в
тело сообщения.

```shell
# fist commit
mkdir docs
touch ./docs/DEVELOPMENT.md

git status
git add .
git commit -m "feat: added blank development guide"
git push --set-upstream origin feature/gif-flow

# second commit
echo "# Руководство для разработчиков" > ./docs/DEVELOPMENT.md

git status
git add .
git commit -m "feat: wrote the development guide"
git push
```

### Pull Request

Когда работа выполнена нужно создать [Pull Request](https://github.com/knowdateapp/knowdateapp/pulls). Pull Request
должен иметь название, отражающее внесенные в код изменения. Если изменения большие или затрагивают несколько
компонентов, то нужно описать их в теле Pull Request'a. При создании Pull Request'а нужно указать проверяющих и связать
его с проектом, в котором ведется трекинг задач, при этом после связывания с проектом нужно указать в связи статус
In Review.

Проверяющий может оставить замечания к добавляемым изменениям или поставить approve. Замечания призывают к обсуждению, а
не беспрекословному их выполнению. Поэтому если вы не согласны с оставленным замечанием, то нужно описать свою точку
зрения и прийти к согласию с проверяющим. Замечание может помечать решенным только его автор после повторной
проверки новых изменений. После закрытия всех замечаний проверяющий должен поставить approve. Pull Request не может быть
влит, если в нем есть открытые замечания или не стоит approve.

### Merge

Перед вливанием ветки `feature/<feature_name>` коммиты в ней должны быть squash'нуты в один коммит. В качестве
инструкции можно ознакомиться со
статьей [Как склеить коммиты и зачем это нужно](https://htmlacademy.ru/blog/git/how-to-squash-commits-and-why-it-is-needed).

```shell
> git cherry -v main   
+ 759b82b75092c965ed1043390af15daa0609ee84 feat: init repository
+ f199c24a7039750fa8a4c2e01aa1e75a3e802276 feat: wrote the development guide
+ ff4b0dd7785a925bf49552a2d1d5005da936f1c9 feat: something else
+ aeb12caa714b78fe4ec2f6f2bfd5f81c862b5e4b feat: added headers to the development guide

> git cherry -v main | wc -l
4

> git rebase -i HEAD~4

> git push -f
```

После этого можно вливать ветку в `main`.

```shell
git checkout main
git merge feature/git-flow
git push
```
