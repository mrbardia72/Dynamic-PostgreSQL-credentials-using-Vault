#!/bin/sh
while true; do
  case "$1" in
    -ud|--up-dev)
        sleep 2
        printf '\033[0;34m ******** deploy dev ********\n \033[0m'
        docker-compose --env-file ../deploy/.env.dev -f ../deploy/docker-compose.yml up --build --detach

        sleep 2
        printf '\033[0;34m ******** configuration vault for secret dynamic database ********\n \033[0m'
        docker exec vault_demo /bin/sh -c "source /config/vault.sh"

        sleep 2
        printf '\033[0;34m ******** migration up db ********\n \033[0m'
        # shellcheck disable=SC2164
        cd ../internal/repository/postgresql/migrations
        sql-migrate up

           exit 0
            ;;
    -dd|--down-dev)
        sleep 2
        printf '\033[0;34m ******** down dev ********\n \033[0m'
        docker-compose --env-file ../deploy/.env.dev -f ../deploy/docker-compose.yml down
      exit 0
      ;;
        -up|--up-prod)
            sleep 2
            printf '\033[0;34m ******** deploy prod ********\n \033[0m'
            docker-compose --env-file ../deploy/.env.prod -f ../deploy/docker-compose.yml up --build --detach

            sleep 2
            printf '\033[0;34m ******** configuration vault for secret dynamic database ********\n \033[0m'
            docker exec vault_demo /bin/sh -c "source /config/vault.sh"

               exit 0
                ;;
        -dp|--down-prod)
            sleep 2
            printf '\033[0;34m ******** down prod ********\n \033[0m'
            docker-compose --env-file ../deploy/.env.prod -f ../deploy/docker-compose.yml down

               exit 0
                ;;
    --)
      break;;
     *)
      printf "sub command usage:  [-ud | --up-dev] [-up | --up-prod] [-dd | --down-dev] [-dp | --down-prod] \n"
      exit 1;;
  esac
done
#Improve your bash/sh shell script with ShellCheck lint script analysis tool

#----------------------------------------------------------------------
#Install ShellCheck on a Debian/Ubuntu Linux
#sudo apt install shellcheck

#----------------------------------------------------------------------
#Install ShellCheck on a CentOS/RHEL/Fedora/Oracle Linux
#First enable EPEL repo on a CentOS/RHEL:
#sudo yum -y install epel-release
#Next, type the following yum command:
#sudo yum install ShellCheck

#----------------------------------------------------------------------
#using a Fedora Linux
#sudo dnf install ShellCheck

#----------------------------------------------------------------------
#Install ShellCheck on an Arch Linux
#sudo pacman -S shellcheck

#----------------------------------------------------------------------
#macOS Unix install ShellCheck
#brew install shellcheck

#----------------------------------------------------------------------
#Run shellcheck deploy.sh in the terminal:
#-- shellcheck deploy.sh  --