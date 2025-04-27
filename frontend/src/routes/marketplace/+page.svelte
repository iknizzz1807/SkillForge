<script lang="ts">
  import type { PageData } from "./$types";
  import { browser } from "$app/environment";
  import { onMount, onDestroy } from "svelte";

  let { data }: { data: PageData } = $props();

  // Đảm bảo skills là mảng
  let projectsDisplay = $state(
    data.projects.map((project) => ({
      ...project,
      skills:
        typeof project.skills === "string"
          ? (project.skills as string).split(",").map((s) => s.trim())
          : project.skills || [],
    }))
  );

  let projectsSuggest: any[] = $state([]);

  // let projectsSuggest = $state(
  //   data.projects
  //     .map((project) => ({
  //       ...project,
  //       skills:
  //         typeof project.skills === "string"
  //           ? (project.skills as string).split(",").map((s) => s.trim())
  //           : project.skills || [],
  //     }))
  //     .slice(0, 5)
  // );

  onMount(async () => {
    if (data.role !== "student") return;
    try {
      const response = await fetch("/api/projects/matching");
      if (response.ok) {
        projectsSuggest = await response.json();
        projectsSuggest = projectsSuggest.map((project) => ({
          ...project,
          // Standardize field names
          id: project.ProjectID || project.project_id,
          title: project.ProjectTitle || project.project_title,
          skills: project.ProjectSkills || project.project_skills || [],
          start_time: project.StartTime || project.start_time,
          end_time: project.EndTime || project.end_time,
          created_by_name: project.CreatorName || project.creator_name,
          match_score: project.match_score,
        }));
      }
    } catch (err) {
      console.log(err);
    }
  });

  // Format date function
  function formatDate(dateString: string): string {
    if (!dateString) return "N/A";

    const date = new Date(dateString);

    // Check if date is valid
    if (isNaN(date.getTime())) return "Invalid date";

    // Format: "Mar 24, 2025"
    return date.toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  }

  // Danh sách các skill categories - dùng để phân loại kỹ năng cho người dùng dễ tìm
  const skillCategories = {
    "Frontend Core": [
      "JavaScript",
      "TypeScript",
      "HTML/CSS",
      "React",
      "Vue",
      "Angular",
      "Svelte",
      "jQuery",
    ],
    "Frontend Frameworks": [
      "Next.js",
      "Nuxt.js",
      "Remix",
      "Gatsby",
      "Preact",
      "SolidJS",
      "Alpine.js",
    ],
    "UI & Styling": [
      "Tailwind CSS",
      "Bootstrap",
      "Material UI",
      "Chakra UI",
      "Ant Design",
      "Styled Components",
      "SASS/SCSS",
      "CSS Modules",
      "CSS-in-JS",
    ],
    "Frontend State & Data": [
      "Redux",
      "MobX",
      "Zustand",
      "Recoil",
      "Jotai",
      "XState",
      "TanStack Query (React Query)",
      "SWR",
      "Apollo Client",
    ],
    "Frontend Build & Tools": [
      "Webpack",
      "Vite",
      "Parcel",
      "Rollup",
      "ESBuild",
      "Babel",
      "ESLint",
      "Prettier",
      "PWA",
      "Storybook",
      "WebComponents",
    ],
    "Frontend Testing": [
      "Jest",
      "Vitest",
      "Testing Library",
      "Cypress",
      "Playwright",
      "Puppeteer",
      "Storybook Testing",
    ],

    "Backend Languages": [
      "Node.js",
      "Python",
      "Java",
      "Kotlin",
      "C#",
      "PHP",
      "Ruby",
      "Go",
      "Rust",
      "Scala",
      "Deno",
    ],
    "Backend Frameworks": [
      "Express",
      "NestJS",
      "Fastify",
      "Koa",
      "Django",
      "Flask",
      "FastAPI",
      "Spring",
      "Spring Boot",
      "ASP.NET",
      ".NET Core",
      "Laravel",
      "Symfony",
      "Ruby on Rails",
      "Gin",
      "Echo",
      "Fiber",
      "Actix",
    ],
    "API Development": [
      "REST API",
      "GraphQL",
      "gRPC",
      "WebSockets",
      "Socket.IO",
      "SignalR",
      "API Gateway",
      "API Documentation",
      "Swagger/OpenAPI",
      "tRPC",
      "JSON:API",
    ],
    "Backend Architecture": [
      "Microservices",
      "Serverless",
      "Monolithic",
      "Event-Driven",
      "CQRS",
      "DDD",
      "Clean Architecture",
      "Hexagonal Architecture",
      "Service-Oriented Architecture",
    ],
    "Authentication & Security": [
      "OAuth",
      "JWT",
      "SAML",
      "OpenID Connect",
      "Two-Factor Authentication",
      "Role-Based Access Control",
      "API Security",
      "Authentication Flows",
      "SSO",
    ],

    "SQL Databases": [
      "SQL",
      "MySQL",
      "PostgreSQL",
      "MariaDB",
      "SQLite",
      "Oracle",
      "SQL Server",
      "Google Cloud SQL",
      "Amazon RDS",
      "Azure SQL Database",
    ],
    "NoSQL Databases": [
      "MongoDB",
      "Cassandra",
      "DynamoDB",
      "Couchbase",
      "Neo4j",
      "Firebase Firestore",
      "Redis",
      "Elasticsearch",
      "CouchDB",
      "RavenDB",
    ],
    "Database Tools & ORM": [
      "Prisma",
      "TypeORM",
      "Sequelize",
      "Mongoose",
      "Knex.js",
      "SQLAlchemy",
      "Entity Framework",
      "Dapper",
      "JDBC",
      "JPA",
      "Hibernate",
    ],
    "Database Administration": [
      "Database Design",
      "Indexing",
      "Query Optimization",
      "Database Migration",
      "Backups & Recovery",
      "Replication",
      "Sharding",
      "Connection Pooling",
      "Database Security",
    ],
    "Modern Data Solutions": [
      "Supabase",
      "PlanetScale",
      "CockroachDB",
      "TiDB",
      "InfluxDB",
      "TimescaleDB",
      "QuestDB",
      "Snowflake",
      "BigQuery",
      "Redshift",
      "Data Warehousing",
    ],

    "DevOps Fundamentals": [
      "CI/CD",
      "Docker",
      "Kubernetes",
      "Infrastructure as Code",
      "Monitoring",
      "Logging",
      "DevSecOps",
      "GitOps",
      "Continuous Deployment",
      "Release Management",
    ],
    "Cloud Platforms - AWS": [
      "AWS",
      "EC2",
      "S3",
      "RDS",
      "Lambda",
      "ECS",
      "EKS",
      "CloudFormation",
      "DynamoDB",
      "AWS CDK",
      "AWS CloudFront",
      "AWS SQS/SNS",
    ],
    "Cloud Platforms - Others": [
      "Azure",
      "Google Cloud",
      "DigitalOcean",
      "Heroku",
      "Vercel",
      "Netlify",
      "Cloudflare",
      "IBM Cloud",
      "Oracle Cloud",
      "Linode",
      "Render",
    ],
    "Infrastructure & Configuration": [
      "Terraform",
      "Pulumi",
      "Ansible",
      "Puppet",
      "Chef",
      "CloudFormation",
      "ARM Templates",
      "Vagrant",
      "Packer",
      "Configuration Management",
    ],
    "CI/CD & Automation": [
      "GitHub Actions",
      "Jenkins",
      "GitLab CI",
      "CircleCI",
      "Travis CI",
      "ArgoCD",
      "Flux",
      "TeamCity",
      "Drone CI",
      "Tekton",
      "Spinnaker",
    ],
    "Monitoring & Observability": [
      "Prometheus",
      "Grafana",
      "ELK Stack",
      "Datadog",
      "New Relic",
      "Sentry",
      "Splunk",
      "Jaeger",
      "OpenTelemetry",
      "Logging",
      "APM",
      "Metrics",
    ],
    "DevOps Tools & Networking": [
      "Linux",
      "Bash/Shell",
      "PowerShell",
      "Nginx",
      "Apache",
      "Istio",
      "Linkerd",
      "Envoy",
      "Load Balancing",
      "HAProxy",
      "CDN",
      "DNS Management",
    ],

    "Data Science Fundamentals": [
      "Python for Data Science",
      "R",
      "Data Analysis",
      "Statistical Analysis",
      "Data Visualization",
      "Exploratory Data Analysis",
      "Data Cleaning",
      "Feature Engineering",
      "Experimental Design",
    ],
    "Machine Learning": [
      "Machine Learning",
      "Supervised Learning",
      "Unsupervised Learning",
      "Classification",
      "Regression",
      "Clustering",
      "Dimensionality Reduction",
      "Ensemble Methods",
      "Feature Selection",
    ],
    "Deep Learning": [
      "Deep Learning",
      "Neural Networks",
      "CNN",
      "RNN",
      "LSTM",
      "Transformers",
      "GAN",
      "Transfer Learning",
      "Fine-tuning",
      "Reinforcement Learning",
    ],
    "ML/AI Frameworks & Libraries": [
      "TensorFlow",
      "PyTorch",
      "Keras",
      "scikit-learn",
      "Hugging Face",
      "LangChain",
      "JAX",
      "MXNet",
      "XGBoost",
      "LightGBM",
      "ONNX",
    ],
    "Data Science Tools": [
      "Pandas",
      "NumPy",
      "SciPy",
      "Matplotlib",
      "Seaborn",
      "Plotly",
      "Jupyter",
      "Dask",
      "Spark",
      "PySpark",
      "Bokeh",
      "Streamlit",
    ],
    "AI Specializations": [
      "Natural Language Processing",
      "Computer Vision",
      "Speech Recognition",
      "Generative AI",
      "Recommendation Systems",
      "Time Series Analysis",
      "Anomaly Detection",
      "AI Ethics",
    ],
    "MLOps & Productionization": [
      "MLOps",
      "ML Pipelines",
      "Model Serving",
      "Model Monitoring",
      "Feature Stores",
      "ML Versioning",
      "ML Testing",
      "Weights & Biases",
      "MLflow",
      "Kubeflow",
      "BentoML",
    ],
    "Big Data": [
      "Big Data",
      "Hadoop",
      "Spark",
      "Kafka",
      "Hive",
      "HBase",
      "Storm",
      "Flink",
      "Beam",
      "Data Lake",
      "ETL",
      "Distributed Computing",
    ],
    "Data Engineering": [
      "Data Engineering",
      "Airflow",
      "dbt",
      "Dagster",
      "Prefect",
      "Fivetran",
      "Stitch",
      "Data Pipelines",
      "Data Warehousing",
      "Data Modeling",
      "Batch Processing",
      "Stream Processing",
    ],

    "Mobile - iOS": [
      "iOS Development",
      "Swift",
      "Objective-C",
      "SwiftUI",
      "UIKit",
      "Core Data",
      "Core Animation",
      "ARKit",
      "WidgetKit",
      "TestFlight",
      "Xcode",
      "CocoaPods",
      "Swift Package Manager",
    ],
    "Mobile - Android": [
      "Android Development",
      "Kotlin for Android",
      "Java for Android",
      "Jetpack Compose",
      "Android SDK",
      "Android Jetpack",
      "Room",
      "WorkManager",
      "Navigation Component",
      "Material Design",
      "Gradle",
    ],
    "Cross-Platform Mobile": [
      "React Native",
      "Flutter",
      "Xamarin",
      "Ionic",
      "Capacitor",
      "NativeScript",
      "Cordova",
      "PWA for Mobile",
      "KMM (Kotlin Multiplatform Mobile)",
      "Unity for Mobile",
    ],
    "Mobile DevOps & QA": [
      "Mobile CI/CD",
      "Fastlane",
      "AppCenter",
      "Mobile Testing",
      "Appium",
      "Espresso",
      "XCTest",
      "Firebase Test Lab",
      "Mobile Analytics",
      "Crash Reporting",
      "Performance Monitoring",
    ],
    "Mobile Features & Services": [
      "Push Notifications",
      "Mobile Authentication",
      "In-App Purchases",
      "Mobile Maps",
      "Geolocation",
      "Bluetooth/BLE",
      "NFC",
      "Camera & Photos",
      "Offline Storage",
      "Background Processing",
    ],

    "UI/UX Design": [
      "UI Design",
      "UX Design",
      "Interaction Design",
      "Visual Design",
      "User Research",
      "Usability Testing",
      "Information Architecture",
      "Wireframing",
      "Prototyping",
      "Design Systems",
      "Responsive Design",
    ],
    "Design Tools": [
      "Figma",
      "Adobe XD",
      "Sketch",
      "InVision",
      "Zeplin",
      "Framer",
      "Photoshop",
      "Illustrator",
      "After Effects",
      "Principle",
      "ProtoPie",
    ],
    "Web Design Specialties": [
      "Web Animation",
      "Microinteractions",
      "Typography",
      "Color Theory",
      "Accessibility (a11y)",
      "Dark Mode Design",
      "Design for Performance",
      "Mobile-First Design",
      "Landing Page Design",
    ],
    "Product Design": [
      "Product Design",
      "Design Thinking",
      "User Personas",
      "User Journeys",
      "A/B Testing",
      "Design Documentation",
      "Design QA",
      "Product Strategy",
      "Brand Design",
      "Content Design",
    ],

    "Blockchain & Crypto": [
      "Blockchain",
      "Bitcoin",
      "Ethereum",
      "Solidity",
      "Web3.js",
      "ethers.js",
      "Smart Contracts",
      "Crypto Wallets",
      "Tokenomics",
      "Consensus Algorithms",
      "Private Blockchains",
    ],
    "Web3 Development": [
      "DApp Development",
      "NFT Development",
      "DeFi",
      "DAOs",
      "Hardhat",
      "Truffle",
      "Foundry",
      "IPFS",
      "The Graph",
      "ERC Standards",
      "Smart Contract Security",
      "Gas Optimization",
    ],
    "Alternative Blockchains": [
      "Solana",
      "Rust for Blockchain",
      "Polkadot",
      "Binance Smart Chain",
      "Avalanche",
      "Polygon",
      "zkRollups",
      "Layer 2 Solutions",
      "Cross-chain Development",
      "Interoperability",
    ],

    "Game Development Core": [
      "Game Design",
      "Unity",
      "Unreal Engine",
      "Godot",
      "C++ for Games",
      "Game Physics",
      "Game AI",
      "Level Design",
      "Game Mechanics",
      "Game Optimization",
      "Game Audio",
    ],
    "Game Art & Animation": [
      "3D Modeling",
      "2D Art",
      "Game Animation",
      "Character Design",
      "Environmental Design",
      "Texturing",
      "Rigging",
      "VFX",
      "Technical Art",
      "Blender",
      "Maya",
    ],
    "Game Tech & Platforms": [
      "Console Development",
      "PC Game Development",
      "Mobile Game Development",
      "WebGL",
      "HTML5 Games",
      "Shader Programming",
      "Multiplayer Networking",
      "Game Backend Services",
      "Game Analytics",
    ],

    "AR/VR/XR": [
      "AR Development",
      "VR Development",
      "MR/XR Development",
      "ARKit",
      "ARCore",
      "Unity XR",
      "WebXR",
      "Three.js",
      "A-Frame",
      "360° Media",
      "Spatial Computing",
      "3D User Interfaces",
    ],
    "AR/VR Hardware": [
      "Meta Quest Development",
      "HoloLens",
      "Apple Vision Pro",
      "Vive",
      "Windows Mixed Reality",
      "OpenXR",
      "AR Glasses",
      "Spatial Audio",
      "Haptic Feedback",
      "Motion Tracking",
    ],

    Cybersecurity: [
      "Cybersecurity",
      "Network Security",
      "Application Security",
      "Cloud Security",
      "Security Architecture",
      "Security Compliance",
      "Threat Modeling",
      "Risk Assessment",
      "Identity Management",
      "Encryption",
    ],
    "Security Testing & Hacking": [
      "Ethical Hacking",
      "Penetration Testing",
      "Security Auditing",
      "Vulnerability Scanning",
      "SAST",
      "DAST",
      "OWASP",
      "Bug Bounty",
      "Red Team",
      "Blue Team",
      "Forensics",
    ],
    "Security Implementation": [
      "Secure Coding",
      "Authentication Implementation",
      "Authorization",
      "API Security",
      "Web Security",
      "Mobile Security",
      "DevSecOps",
      "Container Security",
      "Kubernetes Security",
      "Key Management",
    ],

    "IoT & Embedded": [
      "IoT Development",
      "Embedded Systems",
      "Arduino",
      "Raspberry Pi",
      "ESP32",
      "Microcontrollers",
      "MQTT",
      "CoAP",
      "Embedded C/C++",
      "Embedded Rust",
      "RTOS",
      "Firmware Development",
    ],
    "IoT Connectivity & Ecosystem": [
      "Bluetooth/BLE",
      "Zigbee",
      "LoRaWAN",
      "Z-Wave",
      "Thread",
      "Matter",
      "5G IoT",
      "Sensors & Actuators",
      "Edge Computing",
      "IoT Cloud Platforms",
      "IoT Analytics",
      "IoT Security",
      "Digital Twins",
    ],

    "Project & Team Management": [
      "Agile",
      "Scrum",
      "Kanban",
      "Project Management",
      "Product Management",
      "Team Leadership",
      "Technical Leadership",
      "Stakeholder Management",
      "OKRs",
      "JIRA",
      "Asana",
      "Trello",
      "ClickUp",
    ],
    "Software Craftsmanship": [
      "Clean Code",
      "SOLID Principles",
      "Design Patterns",
      "Test-Driven Development",
      "Refactoring",
      "Code Review",
      "Technical Debt Management",
      "Documentation",
      "Software Architecture",
      "System Design",
    ],
    "Version Control & Collaboration": [
      "Git",
      "GitHub",
      "GitLab",
      "Bitbucket",
      "Collaborative Coding",
      "Trunk-Based Development",
      "Git Flow",
      "GitHub Flow",
      "Conventional Commits",
      "Semantic Versioning",
    ],

    "Career & Soft Skills": [
      "Problem Solving",
      "Analytical Thinking",
      "Communication",
      "Technical Presentations",
      "Technical Writing",
      "Mentoring",
      "Open Source Contribution",
      "Continuous Learning",
      "Interviewing",
      "Career Development",
      "Technical Leadership",
    ],
  };

  // Filter states
  let searchQuery = $state("");
  let selectedDifficulty = $state("All Levels");
  let selectedSkills = $state<string[]>([]);
  let isSkillDropdownOpen = $state(false);
  let skillSearchTerm = $state("");
  let filteredProjects = $state([...projectsDisplay]);

  // Danh sách các mức độ khó
  const difficulties = ["All Levels", "Beginner", "Intermediate", "Expert"];

  // Toggle skill selection
  function toggleSkill(skill: string) {
    if (selectedSkills.includes(skill)) {
      selectedSkills = selectedSkills.filter((s) => s !== skill);
    } else {
      selectedSkills = [...selectedSkills, skill];
    }
  }

  // Close dropdown when clicking outside
  function handleClickOutside(event: MouseEvent) {
    const dropdown = document.getElementById("skills-dropdown");
    const button = document.querySelector(".skills-button");

    if (
      dropdown &&
      !dropdown.contains(event.target as Node) &&
      button &&
      !button.contains(event.target as Node)
    ) {
      isSkillDropdownOpen = false;
    }
  }

  // Apply filters - dùng $effect để theo dõi thay đổi và cập nhật filteredProjects
  $effect(() => {
    filteredProjects = projectsDisplay.filter((project) => {
      // Filter by search query
      const matchesSearch = searchQuery
        ? project.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
          project.description?.toLowerCase().includes(searchQuery.toLowerCase())
        : true;

      // Filter by difficulty
      const matchesDifficulty =
        selectedDifficulty === "All Levels"
          ? true
          : project.difficulty.toLocaleLowerCase() ===
            selectedDifficulty.toLocaleLowerCase();

      // Filter by skills
      const matchesSkills =
        selectedSkills.length === 0
          ? true
          : selectedSkills.some(
              (skill) => project.skills && project.skills.includes(skill)
            );

      return matchesSearch && matchesDifficulty && matchesSkills;
    });
  });

  // Reset filters
  function resetFilters() {
    searchQuery = "";
    selectedDifficulty = "All Levels";
    selectedSkills = [];
    skillSearchTerm = "";
  }

  // Add event listener for click outside
  onMount(() => {
    if (browser) {
      document.addEventListener("click", handleClickOutside);
    }
  });

  onDestroy(() => {
    if (browser) {
      document.removeEventListener("click", handleClickOutside);
    }
  });
</script>

<div class="flex space-x-4 ml-64 pl-4 pr-4 pt-4">
  <!-- Left Column - Contains Filter Panel -->
  <div class="w-2/5 space-y-4 flex-shrink-0">
    <!-- Filter Panel -->
    <div class="card p-4">
      <h3 class="text-base font-semibold mb-3">Filters</h3>
      <div class="space-y-3">
        <!-- Search Projects -->
        <div>
          <label class="block text-sm font-medium mb-1">Search Projects</label>
          <div class="flex items-center space-x-2">
            <input
              type="text"
              bind:value={searchQuery}
              class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
              placeholder="Search by name or description..."
            />
          </div>
        </div>

        <!-- Difficulty -->
        <div>
          <label class="block text-sm font-medium mb-1">Difficulty</label>
          <select
            bind:value={selectedDifficulty}
            class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
          >
            {#each difficulties as difficulty}
              <option value={difficulty}>{difficulty}</option>
            {/each}
          </select>
        </div>

        <!-- Skills with dropdown -->
        <div class="relative">
          <label class="block text-sm font-medium mb-1">Skills</label>
          <div class="relative">
            <div
              class="p-2 border border-gray-300 rounded flex flex-wrap gap-1 min-h-[42px] cursor-pointer focus:outline-none focus:ring-2 focus:ring-[#6b48ff] skills-button"
              onclick={() => (isSkillDropdownOpen = !isSkillDropdownOpen)}
            >
              {#if selectedSkills.length > 0}
                {#each selectedSkills as skill}
                  <span
                    class="bg-[#6b48ff] text-white px-2 py-0.5 text-xs rounded-full flex items-center"
                  >
                    {skill}
                    <button class="ml-1" onclick={() => toggleSkill(skill)}>
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-3 w-3"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                      >
                        <path
                          fill-rule="evenodd"
                          d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                          clip-rule="evenodd"
                        />
                      </svg>
                    </button>
                  </span>
                {/each}
              {:else}
                <span class="text-gray-500 text-sm">Select skills...</span>
              {/if}
            </div>
          </div>

          <!-- Dropdown (đặt bên ngoài thẻ div.relative) -->
          {#if isSkillDropdownOpen}
            <div
              id="skills-dropdown"
              class="absolute z-50 mt-1 w-full bg-white border border-gray-300 rounded shadow-lg max-h-[300px] overflow-y-auto"
            >
              <div class="p-2 border-b sticky top-0 bg-white z-10">
                <input
                  type="text"
                  bind:value={skillSearchTerm}
                  class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-[#6b48ff]"
                  placeholder="Search skills..."
                />
              </div>

              <div class="p-2">
                {#each Object.entries(skillCategories) as [category, categorySkills]}
                  <!-- Chỉ hiển thị lĩnh vực có ít nhất một kỹ năng khớp với từ khóa tìm kiếm -->
                  {@const matchedSkills = categorySkills.filter(
                    (skill) =>
                      !skillSearchTerm ||
                      skill
                        .toLowerCase()
                        .includes(skillSearchTerm.toLowerCase())
                  )}

                  {#if !skillSearchTerm || matchedSkills.length > 0 || category
                      .toLowerCase()
                      .includes(skillSearchTerm.toLowerCase())}
                    <div class="mb-3">
                      <h4
                        class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-1"
                      >
                        {category}
                      </h4>
                      <div class="space-y-1">
                        {#each matchedSkills as skill}
                          <label
                            class="flex items-center space-x-2 p-1 hover:bg-gray-50 rounded cursor-pointer"
                          >
                            <input
                              type="checkbox"
                              checked={selectedSkills.includes(skill)}
                              class="rounded text-[#6b48ff] focus:ring-[#6b48ff]"
                              onchange={() => toggleSkill(skill)}
                            />
                            <span class="text-sm">{skill}</span>
                          </label>
                        {/each}
                      </div>
                    </div>
                  {/if}
                {/each}
              </div>
            </div>
          {/if}
        </div>

        <!-- Action buttons -->
        <div class="flex space-x-2">
          <button
            type="button"
            class="btn-secondary w-full"
            onclick={resetFilters}>Reset Filters</button
          >
        </div>
      </div>
    </div>

    <!-- Suggested Projects -->
    {#if data.role === "student"}
      <div class="card p-3">
        <h3 class="text-base font-semibold mb-2">Suggested Projects</h3>
        <div class="space-y-2">
          {#if projectsSuggest.length === 0}
            <div class="flex justify-center items-center py-4">
              <div
                class="animate-spin rounded-full h-6 w-6 border-b-2 border-[#6b48ff]"
              ></div>
              <span class="ml-2 text-sm text-gray-600"
                >Loading suggestions...</span
              >
            </div>
          {:else}
            {#each projectsSuggest as project}
              <div
                class="flex flex-col p-2 bg-gray-100 rounded hover:bg-gray-200 transition-colors"
              >
                <!-- Match score ở đầu card với chỉ dẫn rõ ràng hơn -->
                <div class="flex justify-start mb-2">
                  <span
                    class="text-xs bg-[#6b48ff] text-white rounded-full px-2 py-0.5 flex items-center"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3 w-3 mr-1"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
                        clip-rule="evenodd"
                      />
                    </svg>
                    <span class="font-bold"
                      >{Math.round(project.match_score)}%</span
                    >
                    <span class="ml-1">skill match</span>
                  </span>
                </div>

                <div class="flex justify-between items-start">
                  <div class="flex-grow">
                    <p class="text-sm font-medium">{project.title}</p>
                    <p class="text-xs text-gray-500 mt-1">
                      Skills: {project.skills.join(", ")}
                    </p>
                    <p class="text-xs text-gray-500">
                      Timeline: {formatDate(project.start_time.toString())} - {formatDate(
                        project.end_time.toString()
                      )} | By
                      <a
                        href="/marketplace"
                        class="text-xs text-[#6b48ff] hover:underline"
                      >
                        {project.created_by_name}
                      </a>
                    </p>
                  </div>
                  <a
                    href={"/marketplace/" + project.id}
                    class="btn text-xs ml-2 flex-shrink-0 self-start">View</a
                  >
                </div>
              </div>
            {/each}
          {/if}
        </div>
      </div>
    {/if}
  </div>

  <!-- Project List -->
  <div class="w-3/5 space-y-3">
    <!-- Projects Count -->
    <div class="flex justify-between items-center">
      <h2 class="text-xl font-semibold">
        Projects ({filteredProjects.length})
      </h2>
    </div>

    <!-- Project Cards -->
    {#if filteredProjects.length === 0}
      <div class="card p-6 text-center">
        <svg
          class="w-16 h-16 text-gray-300 mx-auto mb-4"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
          ></path>
        </svg>
        <h3 class="text-lg font-medium text-gray-700 mb-2">
          No projects found
        </h3>
        <p class="text-gray-500 mb-4">
          Try adjusting your search or filters to find what you're looking for
        </p>
        <button class="btn-secondary" onclick={resetFilters}
          >Clear Filters</button
        >
      </div>
    {:else}
      {#each filteredProjects as project}
        <div class="card p-4 hover:shadow-md transition-shadow">
          <a href={"/marketplace/" + project.id} class="block">
            <div class="flex justify-between items-start">
              <div>
                <h4 class="text-lg font-medium">{project.title}</h4>

                <!-- Difficulty badge -->
                {#if project.difficulty}
                  <span
                    class="inline-block mt-1 mb-2 bg-blue-100 text-blue-800 text-xs px-2 py-0.5 rounded-full"
                    >{project.difficulty}</span
                  >
                {/if}

                <!-- Description -->
                <p class="text-sm text-gray-600 mt-2 line-clamp-2">
                  {project.description || "No description provided."}
                </p>

                <!-- Skills -->
                {#if project.skills && project.skills.length > 0}
                  <div class="flex flex-wrap gap-1 mt-2">
                    {#each project.skills.slice(0, 5) as skill}
                      <span
                        class="bg-gray-100 text-gray-700 text-xs px-2 py-0.5 rounded"
                        >{skill}</span
                      >
                    {/each}
                    {#if project.skills.length > 5}
                      <span
                        class="bg-gray-100 text-gray-700 text-xs px-2 py-0.5 rounded"
                        >+{project.skills.length - 5} more</span
                      >
                    {/if}
                  </div>
                {/if}

                <!-- Timeline and Posted date -->
                <div class="flex flex-wrap text-xs text-gray-400 mt-2">
                  <div class="flex items-center mr-3">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3.5 w-3.5 mr-1"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
                        clip-rule="evenodd"
                      />
                    </svg>
                    <span
                      >Timeline: {formatDate(project.start_time.toString())} - {formatDate(
                        project.end_time.toString()
                      )}</span
                    >
                  </div>
                  <div class="flex items-center">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-3.5 w-3.5 mr-1"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z"
                        clip-rule="evenodd"
                      />
                    </svg>
                    <span
                      >Posted: {formatDate(project.created_at.toString())}</span
                    >
                  </div>
                </div>

                <!-- Creator info -->
                <div class="text-xs text-gray-400 mt-1 flex items-center">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-3.5 w-3.5 mr-1"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
                      clip-rule="evenodd"
                    />
                  </svg>
                  <span>
                    By <span role="link" tabindex="0" class="text-[#6b48ff]"
                      >{project.created_by_name}</span
                    >
                  </span>
                </div>
              </div>

              <div class="flex flex-col items-end space-y-2">
                <!-- Status badge -->
                <span
                  class={`text-xs px-2 py-1 rounded font-bold ${
                    project.status?.toLowerCase() === "open"
                      ? "bg-green-100 text-green-800"
                      : "bg-red-100 text-red-800"
                  }`}
                >
                  {project.status ? project.status.toUpperCase() : "OPEN"}
                </span>

                <!-- Member count -->
                <div class="flex items-center bg-gray-100 px-2 py-1 rounded-lg">
                  <div class="flex -space-x-2 mr-2">
                    <div
                      class="w-6 h-6 rounded-full bg-purple-500 border-2 border-white flex items-center justify-center"
                    >
                      <span class="text-[10px] font-bold text-white"
                        >{project.current_member > 0
                          ? project.current_member
                          : ""}</span
                      >
                    </div>
                    {#if project.current_member >= 2}
                      <div
                        class="w-6 h-6 rounded-full bg-indigo-500 border-2 border-white flex items-center justify-center"
                      ></div>
                    {/if}
                    {#if project.current_member >= 3}
                      <div
                        class="w-6 h-6 rounded-full bg-blue-500 border-2 border-white flex items-center justify-center"
                      ></div>
                    {/if}
                  </div>

                  <div class="flex items-center">
                    <span class="text-xs font-medium">
                      {project.current_member}/{project.max_member}
                    </span>
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 24 24"
                      fill="currentColor"
                      class="w-4 h-4 ml-1 text-gray-500"
                    >
                      <path
                        d="M4.5 6.375a4.125 4.125 0 118.25 0 4.125 4.125 0 01-8.25 0zM14.25 8.625a3.375 3.375 0 116.75 0 3.375 3.375 0 01-6.75 0zM1.5 19.125a7.125 7.125 0 0114.25 0v.003l-.001.119a.75.75 0 01-.363.63 13.067 13.067 0 01-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 01-.364-.63l-.001-.122zM17.25 19.128l-.001.144a2.25 2.25 0 01-.233.96 10.088 10.088 0 005.06-1.01.75.75 0 00.42-.643 4.875 4.875 0 00-6.957-4.611 8.586 8.586 0 011.71 5.157v.003z"
                      />
                    </svg>
                  </div>
                </div>
              </div>
            </div>
          </a>
        </div>
      {/each}
    {/if}
  </div>
</div>
