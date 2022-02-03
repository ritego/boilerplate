git checkout development
git pull --rebase origin development

git checkout staging
git pull --rebase origin staging

git checkout main
git pull --rebase origin main
git checkout development

git checkout main
git merge development
git checkout development
git push origin --all

git checkout staging
git merge development
git checkout development
git push origin --all