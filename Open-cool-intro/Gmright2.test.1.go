2021-03-29T20:51:39.1600299Z ##[section]Starting: Request a runner to run this job
2021-03-29T20:51:39.3271444Z Can't find any online and idle self-hosted runner in current repository that matches the required labels: 'ubuntu-latest'
2021-03-29T20:51:39.3271548Z Can't find any online and idle self-hosted runner in current repository's account/organization that matches the required labels: 'ubuntu-latest'
2021-03-29T20:51:39.3271761Z Found online and idle hosted runner in current repository's account/organization that matches the required labels: 'ubuntu-latest'
2021-03-29T20:51:39.4352126Z ##[section]Finishing: Request a runner to run this job
2021-03-29T20:51:45.0035478Z Current runner version: '2.277.1'
2021-03-29T20:51:45.0071136Z ##[group]Operating System
2021-03-29T20:51:45.0072402Z Ubuntu
2021-03-29T20:51:45.0073020Z 20.04.2
2021-03-29T20:51:45.0073463Z LTS
2021-03-29T20:51:45.0073997Z ##[endgroup]
2021-03-29T20:51:45.0074598Z ##[group]Virtual Environment
2021-03-29T20:51:45.0075294Z Environment: ubuntu-20.04
2021-03-29T20:51:45.0075987Z Version: 20210318.0
2021-03-29T20:51:45.0077056Z Included Software: https://github.com/actions/virtual-environments/blob/ubuntu20/20210318.0/images/linux/Ubuntu2004-README.md
2021-03-29T20:51:45.0078550Z Image Release: https://github.com/actions/virtual-environments/releases/tag/ubuntu20%2F20210318.0
2021-03-29T20:51:45.0079481Z ##[endgroup]
2021-03-29T20:51:45.0081854Z ##[group]GITHUB_TOKEN Permissions
2021-03-29T20:51:45.0083453Z Actions: write
2021-03-29T20:51:45.0084017Z Checks: write
2021-03-29T20:51:45.0084543Z Contents: write
2021-03-29T20:51:45.0085151Z Deployments: write
2021-03-29T20:51:45.0085900Z Issues: write
2021-03-29T20:51:45.0086571Z Metadata: read
2021-03-29T20:51:45.0087210Z OrganizationPackages: write
2021-03-29T20:51:45.0088006Z Packages: write
2021-03-29T20:51:45.0088608Z PullRequests: write
2021-03-29T20:51:45.0089262Z RepositoryProjects: write
2021-03-29T20:51:45.0089994Z SecurityEvents: write
2021-03-29T20:51:45.0090560Z Statuses: write
2021-03-29T20:51:45.0091574Z ##[endgroup]
2021-03-29T20:51:45.0095479Z Prepare workflow directory
2021-03-29T20:51:45.0898654Z Prepare all required actions
2021-03-29T20:51:45.0912331Z Getting action download info
2021-03-29T20:51:45.4786424Z Download action repository 'actions/checkout@v2'
2021-03-29T20:51:47.3885829Z ##[group]Run actions/checkout@v2
2021-03-29T20:51:47.3886817Z with:
2021-03-29T20:51:47.3887490Z   repository: gmright2-Coller/gmright2-cool
2021-03-29T20:51:47.3888660Z   token: ***
2021-03-29T20:51:47.3889208Z   ssh-strict: true
2021-03-29T20:51:47.3890114Z   persist-credentials: true
2021-03-29T20:51:47.3890735Z   clean: true
2021-03-29T20:51:47.3891511Z   fetch-depth: 1
2021-03-29T20:51:47.3892106Z   lfs: false
2021-03-29T20:51:47.3892717Z   submodules: false
2021-03-29T20:51:47.3893643Z ##[endgroup]
2021-03-29T20:51:48.2101788Z Syncing repository: gmright2-Coller/gmright2-cool
2021-03-29T20:51:48.2103260Z ##[group]Getting Git version info
2021-03-29T20:51:48.2104669Z Working directory is '/home/runner/work/gmright2-cool/gmright2-cool'
2021-03-29T20:51:48.2106253Z [command]/usr/bin/git version
2021-03-29T20:51:48.2106994Z git version 2.31.0
2021-03-29T20:51:48.2108405Z ##[endgroup]
2021-03-29T20:51:48.2109739Z Deleting the contents of '/home/runner/work/gmright2-cool/gmright2-cool'
2021-03-29T20:51:48.2111841Z ##[group]Initializing the repository
2021-03-29T20:51:48.2113340Z [command]/usr/bin/git init /home/runner/work/gmright2-cool/gmright2-cool
2021-03-29T20:51:48.2114787Z hint: Using 'master' as the name for the initial branch. This default branch name
2021-03-29T20:51:48.2115901Z hint: is subject to change. To configure the initial branch name to use in all
2021-03-29T20:51:48.2117303Z hint: of your new repositories, which will suppress this warning, call:
2021-03-29T20:51:48.2118173Z hint: 
2021-03-29T20:51:48.2119112Z hint: 	git config --global init.defaultBranch <name>
2021-03-29T20:51:48.2119766Z hint: 
2021-03-29T20:51:48.2120672Z hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
2021-03-29T20:51:48.2121807Z hint: 'development'. The just-created branch can be renamed via this command:
2021-03-29T20:51:48.2122600Z hint: 
2021-03-29T20:51:48.2123277Z hint: 	git branch -m <name>
2021-03-29T20:51:48.2124647Z Initialized empty Git repository in /home/runner/work/gmright2-cool/gmright2-cool/.git/
2021-03-29T20:51:48.2126066Z [command]/usr/bin/git remote add origin https://github.com/gmright2-Coller/gmright2-cool
2021-03-29T20:51:48.2127036Z ##[endgroup]
2021-03-29T20:51:48.2128127Z ##[group]Disabling automatic garbage collection
2021-03-29T20:51:48.2129346Z [command]/usr/bin/git config --local gc.auto 0
2021-03-29T20:51:48.2130507Z ##[endgroup]
2021-03-29T20:51:48.2133644Z ##[group]Setting up auth
2021-03-29T20:51:48.2134871Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2021-03-29T20:51:48.2136653Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :
2021-03-29T20:51:48.2138309Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2021-03-29T20:51:48.2140509Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :
2021-03-29T20:51:48.2142497Z [command]/usr/bin/git config --local http.https://github.com/.extraheader AUTHORIZATION: basic ***
2021-03-29T20:51:48.2143532Z ##[endgroup]
2021-03-29T20:51:48.2144308Z ##[group]Fetching the repository
2021-03-29T20:51:48.2146633Z [command]/usr/bin/git -c protocol.version=2 fetch --no-tags --prune --progress --no-recurse-submodules --depth=1 origin +588608e0d5d842a4de070f42e29532d785e86343:refs/remotes/origin/master
2021-03-29T20:51:48.4585215Z remote: Enumerating objects: 49, done.        
2021-03-29T20:51:48.4586471Z remote: Counting objects:   2% (1/49)        
2021-03-29T20:51:48.4587369Z remote: Counting objects:   4% (2/49)        
2021-03-29T20:51:48.4588233Z remote: Counting objects:   6% (3/49)        
2021-03-29T20:51:48.4588916Z remote: Counting objects:   8% (4/49)        
2021-03-29T20:51:48.4589619Z remote: Counting objects:  10% (5/49)        
2021-03-29T20:51:48.4590514Z remote: Counting objects:  12% (6/49)        
2021-03-29T20:51:48.4591324Z remote: Counting objects:  14% (7/49)        
2021-03-29T20:51:48.4591972Z remote: Counting objects:  16% (8/49)        
2021-03-29T20:51:48.4592638Z remote: Counting objects:  18% (9/49)        
2021-03-29T20:51:48.4593346Z remote: Counting objects:  20% (10/49)        
2021-03-29T20:51:48.4594042Z remote: Counting objects:  22% (11/49)        
2021-03-29T20:51:48.4594767Z remote: Counting objects:  24% (12/49)        
2021-03-29T20:51:48.4595417Z remote: Counting objects:  26% (13/49)        
2021-03-29T20:51:48.4596128Z remote: Counting objects:  28% (14/49)        
2021-03-29T20:51:48.4596770Z remote: Counting objects:  30% (15/49)        
2021-03-29T20:51:48.4597468Z remote: Counting objects:  32% (16/49)        
2021-03-29T20:51:48.4598108Z remote: Counting objects:  34% (17/49)        
2021-03-29T20:51:48.4598781Z remote: Counting objects:  36% (18/49)        
2021-03-29T20:51:48.4599508Z remote: Counting objects:  38% (19/49)        
2021-03-29T20:51:48.4601100Z remote: Counting objects:  40% (20/49)        
2021-03-29T20:51:48.4601931Z remote: Counting objects:  42% (21/49)        
2021-03-29T20:51:48.4602586Z remote: Counting objects:  44% (22/49)        
2021-03-29T20:51:48.4603368Z remote: Counting objects:  46% (23/49)        
2021-03-29T20:51:48.4604046Z remote: Counting objects:  48% (24/49)        
2021-03-29T20:51:48.4604773Z remote: Counting objects:  51% (25/49)        
2021-03-29T20:51:48.4605414Z remote: Counting objects:  53% (26/49)        
2021-03-29T20:51:48.4606648Z remote: Counting objects:  55% (27/49)        
2021-03-29T20:51:48.4607365Z remote: Counting objects:  57% (28/49)        
2021-03-29T20:51:48.4608071Z remote: Counting objects:  59% (29/49)        
2021-03-29T20:51:48.4608789Z remote: Counting objects:  61% (30/49)        
2021-03-29T20:51:48.4609434Z remote: Counting objects:  63% (31/49)        
2021-03-29T20:51:48.4610150Z remote: Counting objects:  65% (32/49)        
2021-03-29T20:51:48.4610815Z remote: Counting objects:  67% (33/49)        
2021-03-29T20:51:48.4612300Z remote: Counting objects:  69% (34/49)        
2021-03-29T20:51:48.4612970Z remote: Counting objects:  71% (35/49)        
2021-03-29T20:51:48.4613612Z remote: Counting objects:  73% (36/49)        
2021-03-29T20:51:48.4614308Z remote: Counting objects:  75% (37/49)        
2021-03-29T20:51:48.4615208Z remote: Counting objects:  77% (38/49)        
2021-03-29T20:51:48.4615932Z remote: Counting objects:  79% (39/49)        
2021-03-29T20:51:48.4617148Z remote: Counting objects:  81% (40/49)        
2021-03-29T20:51:48.4617889Z remote: Counting objects:  83% (41/49)        
2021-03-29T20:51:48.4618535Z remote: Counting objects:  85% (42/49)        
2021-03-29T20:51:48.4619240Z remote: Counting objects:  87% (43/49)        
2021-03-29T20:51:48.4619886Z remote: Counting objects:  89% (44/49)        
2021-03-29T20:51:48.4620545Z remote: Counting objects:  91% (45/49)        
2021-03-29T20:51:48.4621318Z remote: Counting objects:  93% (46/49)        
2021-03-29T20:51:48.4622419Z remote: Counting objects:  95% (47/49)        
2021-03-29T20:51:48.4623157Z remote: Counting objects:  97% (48/49)        
2021-03-29T20:51:48.4623806Z remote: Counting objects: 100% (49/49)        
2021-03-29T20:51:48.4624633Z remote: Counting objects: 100% (49/49), done.        
2021-03-29T20:51:48.4625382Z remote: Compressing objects:   2% (1/38)        
2021-03-29T20:51:48.4626155Z remote: Compressing objects:   5% (2/38)        
2021-03-29T20:51:48.4626847Z remote: Compressing objects:   7% (3/38)        
2021-03-29T20:51:48.4628089Z remote: Compressing objects:  10% (4/38)        
2021-03-29T20:51:48.4628864Z remote: Compressing objects:  13% (5/38)        
2021-03-29T20:51:48.4629553Z remote: Compressing objects:  15% (6/38)        
2021-03-29T20:51:48.4630350Z remote: Compressing objects:  18% (7/38)        
2021-03-29T20:51:48.4631040Z remote: Compressing objects:  21% (8/38)        
2021-03-29T20:51:48.4632045Z remote: Compressing objects:  23% (9/38)        
2021-03-29T20:51:48.4633425Z remote: Compressing objects:  26% (10/38)        
2021-03-29T20:51:48.4634200Z remote: Compressing objects:  28% (11/38)        
2021-03-29T20:51:48.4634945Z remote: Compressing objects:  31% (12/38)        
2021-03-29T20:51:48.4635905Z remote: Compressing objects:  34% (13/38)        
2021-03-29T20:51:48.4636600Z remote: Compressing objects:  36% (14/38)        
2021-03-29T20:51:48.4665163Z remote: Compressing objects:  39% (15/38)        
2021-03-29T20:51:48.4666660Z remote: Compressing objects:  42% (16/38)        
2021-03-29T20:51:48.4667609Z remote: Compressing objects:  44% (17/38)        
2021-03-29T20:51:48.4668315Z remote: Compressing objects:  47% (18/38)        
2021-03-29T20:51:48.4669062Z remote: Compressing objects:  50% (19/38)        
2021-03-29T20:51:48.4669779Z remote: Compressing objects:  52% (20/38)        
2021-03-29T20:51:48.4670472Z remote: Compressing objects:  55% (21/38)        
2021-03-29T20:51:48.4671815Z remote: Compressing objects:  57% (22/38)        
2021-03-29T20:51:48.4672535Z remote: Compressing objects:  60% (23/38)        
2021-03-29T20:51:48.4673366Z remote: Compressing objects:  63% (24/38)        
2021-03-29T20:51:48.4674101Z remote: Compressing objects:  65% (25/38)        
2021-03-29T20:51:48.4674853Z remote: Compressing objects:  68% (26/38)        
2021-03-29T20:51:48.4675542Z remote: Compressing objects:  71% (27/38)        
2021-03-29T20:51:48.4676338Z remote: Compressing objects:  73% (28/38)        
2021-03-29T20:51:48.4677605Z remote: Compressing objects:  76% (29/38)        
2021-03-29T20:51:48.4678381Z remote: Compressing objects:  78% (30/38)        
2021-03-29T20:51:48.4679160Z remote: Compressing objects:  81% (31/38)        
2021-03-29T20:51:48.4679857Z remote: Compressing objects:  84% (32/38)        
2021-03-29T20:51:48.4680734Z remote: Compressing objects:  86% (33/38)        
2021-03-29T20:51:48.4681872Z remote: Compressing objects:  89% (34/38)        
2021-03-29T20:51:48.4682944Z remote: Compressing objects:  92% (35/38)        
2021-03-29T20:51:48.4683694Z remote: Compressing objects:  94% (36/38)        
2021-03-29T20:51:48.4684999Z remote: Compressing objects:  97% (37/38)        
2021-03-29T20:51:48.4685943Z remote: Compressing objects: 100% (38/38)        
2021-03-29T20:51:48.4686780Z remote: Compressing objects: 100% (38/38), done.        
2021-03-29T20:51:49.5720442Z remote: Total 49 (delta 0), reused 36 (delta 0), pack-reused 0        
2021-03-29T20:51:49.6497506Z From https://github.com/gmright2-Coller/gmright2-cool
2021-03-29T20:51:49.6500057Z  * [new ref]         588608e0d5d842a4de070f42e29532d785e86343 -> origin/master
2021-03-29T20:51:49.6605452Z ##[endgroup]
2021-03-29T20:51:49.6611367Z ##[group]Determining the checkout info
2021-03-29T20:51:49.6617206Z ##[endgroup]
2021-03-29T20:51:49.6619814Z ##[group]Checking out the ref
2021-03-29T20:51:49.6622558Z [command]/usr/bin/git checkout --progress --force -B master refs/remotes/origin/master
2021-03-29T20:51:49.7235435Z Reset branch 'master'
2021-03-29T20:51:49.7237833Z Branch 'master' set up to track remote branch 'master' from 'origin'.
2021-03-29T20:51:49.7239974Z ##[endgroup]
2021-03-29T20:51:49.7286650Z [command]/usr/bin/git log -1 --format='%H'
2021-03-29T20:51:49.7288805Z '588608e0d5d842a4de070f42e29532d785e86343'
2021-03-29T20:51:49.7509415Z ##[group]Run echo Hello, world!
2021-03-29T20:51:49.7510118Z [36;1mecho Hello, world![0m
2021-03-29T20:51:49.7560028Z shell: /usr/bin/bash -e {0}
2021-03-29T20:51:49.7560933Z ##[endgroup]
2021-03-29T20:51:49.7640465Z Hello, world!
2021-03-29T20:51:49.7894264Z ##[group]Run echo Add other actions to build,
2021-03-29T20:51:49.7894995Z [36;1mecho Add other actions to build,[0m
2021-03-29T20:51:49.7895552Z [36;1mecho test, and deploy your project.[0m
2021-03-29T20:51:49.7945176Z shell: /usr/bin/bash -e {0}
2021-03-29T20:51:49.7945770Z ##[endgroup]
2021-03-29T20:51:49.8030709Z Add other actions to build,
2021-03-29T20:51:49.8031292Z test, and deploy your project.
2021-03-29T20:51:49.8115627Z Post job cleanup.
2021-03-29T20:51:49.9302172Z [command]/usr/bin/git version
2021-03-29T20:51:49.9356663Z git version 2.31.0
2021-03-29T20:51:49.9398788Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2021-03-29T20:51:49.9436308Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :
2021-03-29T20:51:49.9702907Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2021-03-29T20:51:49.9735993Z http.https://github.com/.extraheader
2021-03-29T20:51:49.9744750Z [command]/usr/bin/git config --local --unset-all http.https://github.com/.extraheader
2021-03-29T20:51:49.9780416Z [command]/usr/bin/git submodule foreach --recursive git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :
2021-03-29T20:51:50.0091272Z Cleaning up orphan processes
