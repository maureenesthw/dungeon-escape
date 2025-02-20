# **Dungeon Escape (Text Adventure) 🏰**
<div align="center">
  <img src="statics/dungeon-escape.gif">
</div>

## **🛠️ How to Run**
1. Clone this repository
2. Run `go run .` command

## **🛠️ Game Requirements**

### **1. Game Objective 🎯**
- The player is trapped in a dungeon and must escape within **10 moves**.
- The player has **100 HP** and loses if HP reaches **0**.
- The player must **find a key** before they can escape.

---

### **2. Random Events (Triggered After Choosing a Door) 🔄**
Each time the player chooses a door (**Left, Right, or Forward**), they will encounter **one** of these random events:

#### **🗝️ 1. Finding a Key (30% chance)**
- "You found a key! Now you can escape if you find the exit!"

#### **🐉 2. Monster Encounter (35% chance)**
- "A monster appears! Fight (F) or Run (R)?"
- Fighting: **50% chance to win** (no damage), **50% chance to lose** (-30 HP).
- Running: **70% chance to escape**, else **lose 15 HP**.

#### **⚠️ 3. Trap (20% chance)**
- "You stepped on a trap! You lost 20 HP!"

#### **🚪 4. Exit Door (15% chance, only works with key)**
- If the player **has the key**, they escape and **win**. 
- If they **don’t have the key**, they get:
- "The door is locked! Keep searching!"

---

### **3. Player Actions**
- Choose a door: **Left (L), Right (R), or Forward (F)**.
- If they find a monster: **Fight (F) or Run (R)**.
- If they find the exit: **Escape (E) if they have a key**.

---

### **4. Winning & Losing Conditions**
- **Win:** Escape through the exit **with a key** before **10 moves run out**.
- **Lose:** If **HP reaches 0** or **moves run out** before escaping.

---

Try it out 🚀